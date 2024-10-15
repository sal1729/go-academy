package main

import (
	"log"
	"os"
	"path/filepath"
	todo "todo_app_functions"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("todo_app: ")
	log.SetFlags(0)
	// Also set up a function for catching errors
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// TODO Swap out for a keyboard input?
	dir, err := os.Getwd()
	catchError(err)

	filename := filepath.Join(dir, "files/todo-list.json")
	todoListBytes, err := todo.JsonBytesFromFile(filename)
	catchError(err)

	// We don't need to unmarshal if all we want to do is look
	// But note, we're not validating file content here
	jsonErr := todo.PrettyPrintJson(todoListBytes)
	if jsonErr != nil {
		log.Printf("Error parsing json: %s\n", jsonErr)
	}
}
