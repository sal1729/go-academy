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
// Run `go build -o franz_web exercise-18.go` to construct the executable
// Our app is called franz, after Franz Liszt.
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("ToDo: ")
	log.SetFlags(0)

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		log.Fatal(dirErr)
	}
	datasource := todo.LocalDatasource{
		Filepath: filepath.Join(dir, fmt.Sprintf("data/%s", "todo-list.json")),
	}
	http.HandleFunc("/list", makeHandler(listHandler, datasource))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
