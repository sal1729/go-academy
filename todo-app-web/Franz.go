package main

// This stuff mostly from https://go.dev/doc/articles/wiki/

import (
	"fmt"
	"html/template"
	"net/http"
	todo "todo_app_functions"
)

type Page struct {
	Tasks  todo.Entries
	Filter string
}

var templates = template.Must(template.ParseFiles("franz.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request, d todo.Datasource) {
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

	filterString := ""
	if data.Status != "" {
		filterString = fmt.Sprintf("\"%s\" ", data.Status)
	}

	page := Page{
		Tasks:  list,
		Filter: fmt.Sprintf(" Showing all %stasks.", filterString),
	}
	renderTemplate(w, "franz", page)
}

// TODO Put path validation back in?
//var validPath = regexp.MustCompile("^/(list)/([a-zA-Z0-9]+)$")

// A closure allows us to inject the datasource into the handlers, instantiating it only once
func makeHandler(fn func(http.ResponseWriter, *http.Request, todo.Datasource), datasource todo.Datasource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, datasource)
	}
}
