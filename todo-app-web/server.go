package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	todo "todo_app_functions"
)

// Based on https://go.dev/doc/articles/wiki/
// Run `go build -o franz_web server.go` to construct the executable
// Our app is called franz, after Franz Liszt.
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// This is pretty useless - we never pass the logger through anywhere
	log.SetPrefix("ToDo: ")
	log.SetFlags(0)

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	datasource := todo.InitialiseDatasource(filepath.Join(dir, fmt.Sprintf("data/%s", "todo-list.json")))

	http.HandleFunc("/create", makeHandler(createHandler, datasource))
	http.HandleFunc("/list", makeHandler(listHandler, datasource))
	http.HandleFunc("/update", makeHandler(updateHandler, datasource))
	http.HandleFunc("/delete", makeHandler(deleteHandler, datasource))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
