package main

// Started life as todo-app-web/Franz.go

import (
	"encoding/json"
	"fmt"
	franz "franz_brain"
	"html/template"
	"net/http"
)

type Page struct {
	Tasks  []franz.ListItem
	Filter string
}

var templates = template.Must(template.ParseFiles("franz.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// When you hit the app for the first time, or possibly do a hard refresh, we give you everything 🎁
	status := "all"

	// Read tasks according to filter
	queryUrl := fmt.Sprintf("http://localhost:8080/api?status=%s", status)
	resp, err := http.Get(queryUrl)
	if err != nil {
		http.Error(w, "Unable to retrieve tasks", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close() // TODO handle error here

	// Convert the initial list and status into page data
	var list []franz.ListItem
	if err := json.NewDecoder(resp.Body).Decode(&list); err != nil {
		http.Error(w, "Error reading tasks", http.StatusInternalServerError)
		return
	}

	page := Page{
		Tasks:  list,
		Filter: status,
	}
	renderTemplate(w, "franz", page)
}
