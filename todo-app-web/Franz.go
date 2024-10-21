package main

// This stuff mostly from https://go.dev/doc/articles/wiki/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
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
func renderReq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	// Output the HTTP request details to the webpage
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Headers:\n")

	// Loop through the headers and print each one
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}

	// Print body if there's any (for POST/PUT methods)
	fmt.Fprintf(w, "\nBody:\n")
	if r.Body != nil {
		// Read the body
		buf := make([]byte, 1024)
		n, _ := r.Body.Read(buf)
		if n > 0 {
			fmt.Fprintf(w, "%s\n", string(buf[:n]))
		} else {
			fmt.Fprintf(w, "No body")
		}
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request, d todo.Datasource) {
	var data todo.CrudRequest

	switch r.Method {
	case "GET":
		// Extract filter
		fmt.Println(r.URL.String())
		decodedListFilter, err := url.QueryUnescape(r.URL.String())
		if err != nil {
			fmt.Println("Error decoding:", err)
			return
		}
		data.Status = decodedListFilter
		data.Task = "all"

		// Get filtered list, based on request url
		list, listErr := d.Read(data.Task, data.Status)
		if listErr != nil {
			http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, listErr), http.StatusBadRequest)
			return
		}

		// Render page with filters applied
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

	case "POST":
		// Extract body
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			// http.Error(w, "Invalid JSON", http.StatusBadRequest)
			renderReq(w, r)
			return
		}

		// Create new task
		_, createErr := d.Create(data.Task, data.Status)
		if createErr != nil {
			http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, createErr), http.StatusBadRequest)
			return
		}

		// Route back to list page to view changes
		http.Redirect(w, r, "/list", http.StatusFound)

	case "PUT":
		// Extract body
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			// http.Error(w, "Invalid JSON", http.StatusBadRequest)
			renderReq(w, r)
			return
		}

		// Update task
		_, updateErr := d.Update(data.Task, data.Status)
		if updateErr != nil {
			http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, updateErr), http.StatusBadRequest)
			return
		}

		// Route back to list page to view changes
		http.Redirect(w, r, "/list", http.StatusFound)

	case "DELETE":
		// Extract body
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			// http.Error(w, "Invalid JSON", http.StatusBadRequest)
			renderReq(w, r)
			return
		}

		// Delete task(s)
		_, deleteErr := d.Delete(data.Task, data.Status)
		if deleteErr != nil {
			http.Error(w, fmt.Sprintf("Error handling request: %s, Error: %v", data, deleteErr), http.StatusBadRequest)
			return
		}

		// Route back to list page to view changes
		http.Redirect(w, r, "/list", http.StatusFound)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// A closure allows us to inject the datasource into the handlers, instantiating it only once
// The "actionHandler" was a bit of a nasty hack to get around the frontend only using GET and POST requests
// We use the app routing and the form inputs to construct the api request
func actionHandler(action string, d todo.Datasource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody := todo.CrudRequest{
			Action: action,
		}
		var requestMethod string
		requestUrl := "all" // We'll use the url to filter the list on the screen. V. nasty.

		// Clone the original request, we'll modify it below
		newRequest, err := http.NewRequest("", requestUrl, r.Body)
		if err != nil {
			http.Error(w, "Error constructing new request", http.StatusInternalServerError)
			return
		}

		// Copy the form data, etc. from the original request
		newRequest.Form = r.Form

		// Set the correct method and body
		switch action {
		case "list":
			requestMethod = "GET"
			fmt.Println("got to list action") // TODO delete
			if r.Method == http.MethodPost {
				err := r.ParseForm()
				if err != nil {
					http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
					return
				}

				requestUrl = url.QueryEscape(r.FormValue("status"))

			} else if r.Method == http.MethodGet {
				// Set default to list all
				requestUrl = "all"
			} else {
				http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
				return
			}

		case "create":
			requestMethod = "POST"

			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
				return
			}
			requestBody.Task = r.FormValue("task")
			requestBody.Status = r.FormValue("status")

		case "update":
			requestMethod = "PUT"

			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
				return
			}
			requestBody.Task = r.FormValue("task")
			requestBody.Status = r.FormValue("status")

		case "delete":
			requestMethod = "DELETE"

			err := r.ParseForm()
			if err != nil {
				http.Error(w, fmt.Sprintf("Error parsing form: %v", err), http.StatusBadRequest)
				return
			}
			requestBody.Task = r.FormValue("task")
			requestBody.Status = r.FormValue("status")

		default:
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		// Construct the new request
		// Serialize the struct to JSON
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			eString := fmt.Sprintln("Error marshaling JSON:", err)
			http.Error(w, eString, http.StatusBadRequest)
			return
		}

		// Create a new request with the CrudRequest JSON body
		req, err := http.NewRequest(requestMethod, requestUrl, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		// Set the Content-Type header to indicate JSON data
		req.Header.Set("Content-Type", "application/json")

		// Call taskHandler with the newly constructed request
		fmt.Println(newRequest)
		fmt.Println(requestBody)
		taskHandler(w, newRequest, d)
	}
}
