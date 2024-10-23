package franz

import (
	"encoding/json"
	"net/http"
)

// Api expects a POST, PUT or DELETE request with CrudRequest in body, or a GET request which may or may not define a status
// TODO Do error handling in cases where create fails, read task does not exist etc
func Api(w http.ResponseWriter, r *http.Request, requests chan AccessRequest) {
	var crudRequest CrudRequest

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		if err := json.NewDecoder(r.Body).Decode(&crudRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		success := make(chan SuccessResponse)
		requests <- AccessRequest{
			CrudRequest:   crudRequest,
			ActionSuccess: success,
		}

		complete := <-success
		if !complete.Success || complete.Error != nil {
			http.Error(w, "Failed create request", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))

	case http.MethodGet:
		// Through the API we only allow read all and read by status, not read for individual tasks
		status := r.URL.Query().Get("status")
		crudRequest.Action = "read"
		crudRequest.Task = "all"
		crudRequest.Status = status

		success := make(chan SuccessResponse)
		response := make(chan []ListItem)
		requests <- AccessRequest{
			CrudRequest:   crudRequest,
			ActionSuccess: success,
			Result:        response,
		}

		complete := <-success
		if !complete.Success || complete.Error != nil {
			http.Error(w, "Failed create request", http.StatusBadRequest)
			return
		}

		list := <-response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(list)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// MakeHandler is a closure. A closure allows us to inject the datasource request channel into the handlers
func MakeHandler(fn func(http.ResponseWriter, *http.Request, chan AccessRequest), requests chan AccessRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, requests)
	}
}
