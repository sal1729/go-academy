package main

// This stuff mostly from https://go.dev/doc/articles/wiki/

import (
	"fmt"
	"html/template"
	"net/http"
	todo "todo_app_functions"
)

type Page struct {
	Tasks        todo.Entries
	Filter       string
	FilterString string
}

var templates = template.Must(template.ParseFiles("franz.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request, d todo.MutexDatasource) {
	data := todo.CrudRequest{
		Action: "create",
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
			return
		}

		data.Task = r.FormValue("task")
		data.Status = r.FormValue("status")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Handle request
	_, createErr := d.Create(data.Task, data.Status)
	if createErr != nil {
		http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, createErr), http.StatusBadRequest)
		return
	}

	// Redirect back to list page
	http.Redirect(w, r, "/list", http.StatusFound)
}

func listHandler(w http.ResponseWriter, r *http.Request, d todo.MutexDatasource) {
	data := todo.CrudRequest{
		Action: "list",
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
			return
		}

		data.Task = "all"
		if r.FormValue("status") == "All" {
			data.Status = ""
		} else {
			data.Status = r.FormValue("status")
		}

	} else if r.Method == http.MethodGet {
		// Set default request to list all
		data.Task = "all"
		data.Status = ""
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Handle request
	list, listErr := d.Read(data.Task, data.Status)
	if listErr != nil {
		http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, listErr), http.StatusBadRequest)
		return
	}

	filterString := " "
	if data.Status != "" {
		filterString = fmt.Sprintf(" \"%s\" ", data.Status)
	}

	page := Page{
		Tasks:        list,
		Filter:       data.Status,
		FilterString: filterString,
	}
	renderTemplate(w, "franz", page)
}

func updateHandler(w http.ResponseWriter, r *http.Request, d todo.MutexDatasource) {
	data := todo.CrudRequest{ // This block is redundant in each handler as we call the correct action on the datasource
		Action: "update", // This handler is almost identical to the createHandler TODO dedupe?
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
			return
		}

		data.Task = r.FormValue("task")
		data.Status = r.FormValue("status")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Handle request
	_, updateErr := d.Update(data.Task, data.Status)
	if updateErr != nil {
		http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, updateErr), http.StatusBadRequest)
		return
	}

	// Redirect back to list page
	http.Redirect(w, r, "/list", http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, d todo.MutexDatasource) {
	data := todo.CrudRequest{ // This block is redundant in each handler as we call the correct action on the datasource
		Action: "delete", // This handler is almost identical to the createHandler TODO dedupe?
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
			return
		}

		data.Task = r.FormValue("task")
		data.Status = r.FormValue("status")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Handle request
	_, deleteErr := d.Delete(data.Task, data.Status)
	if deleteErr != nil {
		http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, deleteErr), http.StatusBadRequest)
		return
	}

	// Redirect back to list page
	http.Redirect(w, r, "/list", http.StatusFound)
}

// A closure allows us to inject the datasource into the handlers, instantiating it only once
func makeHandler(fn func(http.ResponseWriter, *http.Request, todo.MutexDatasource), datasource todo.MutexDatasource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, datasource)
	}
}
