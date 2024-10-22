package main

import (
	"encoding/json"
	franz "franz_datastore"
	"net/http"
)

// TODO Do error handling in cases where create fails, read task does not exist etc
// franzApi could almost certainly be better. He's not following any standard or requirements other then my own made up nonsense.
func franzApi(w http.ResponseWriter, r *http.Request, todoList *[]franz.ListItem) {
	var crudRequest franz.CrudRequest

	switch r.Method {
	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&crudRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, _ := franz.CreateTask(todoList, crudRequest) // TODO handle erroring here

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)

	case http.MethodGet:
		// Through the API we only allow read all and read by status, not read for individual tasks
		status := r.URL.Query().Get("status")
		// status = "" is permissible => read all
		crudRequest.Action = "read"
		crudRequest.Task = "all"
		crudRequest.Status = status
		result, _ := franz.ReadFromList(todoList, crudRequest)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	case http.MethodPut:
		if err := json.NewDecoder(r.Body).Decode(&crudRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, err := franz.UpdateListItems(todoList, crudRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // update only errors if there's a problem with the requested status
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	case http.MethodDelete:
		if err := json.NewDecoder(r.Body).Decode(&crudRequest); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		result, _ := franz.DeleteFromList(todoList, crudRequest)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// A closure allows us to inject the datasource into the handlers, instantiating it only once
func makeHandler(fn func(http.ResponseWriter, *http.Request, *[]franz.ListItem), data *[]franz.ListItem) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, data)
	}
}
