package main

import (
	"encoding/json"
	franz "franz_datastore"
	"net/http"
)

// TODO Do error handling in cases where create fails, read task does not exist etc
// franzApi expects a POST, PUT or DELETE request with CrudRequest in body, or a GET request which may or may not define a status
func franzApi(w http.ResponseWriter, r *http.Request, requests chan franz.AccessRequest) {
	var crudRequest franz.CrudRequest

	switch r.Method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		if err := json.NewDecoder(r.Body).Decode(&crudRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		success := make(chan franz.SuccessResponse)
		requests <- franz.AccessRequest{
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

		success := make(chan franz.SuccessResponse)
		response := make(chan []franz.ListItem)
		requests <- franz.AccessRequest{
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

// A closure allows us to inject the datasource request channel into the handlers
func makeHandler(fn func(http.ResponseWriter, *http.Request, chan franz.AccessRequest), requests chan franz.AccessRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, requests)
	}
}
