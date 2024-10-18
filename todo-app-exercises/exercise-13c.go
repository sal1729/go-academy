package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	todo "todo_app_functions"
)

// exercise-13 asked us for a cli tool...
// Run `go build -o todo-cli exercise-13c.go` to construct the executable
// Then `./todo-cli -filename=todo-list-short.json -tasks=12`, to run, for example
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("readTodoFromFile: ")
	log.SetFlags(0)
	// Also set up a function for catching errors
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Option flags
	filename := flag.String("filename", "todo-list.json", "filename to read")
	tasksToRead := flag.Int("tasks", 10, "number of tasks to read")
	flag.Parse()

	dir, dirErr := os.Getwd()
	catchError(dirErr)
	file := filepath.Join(dir, fmt.Sprintf("data/%s", *filename))

	todoListEntries, err := todo.StreamEntriesFromFile(file, *tasksToRead)
	catchError(err)

	log.Printf("The first %d tasks are", len(todoListEntries))
	todo.PrintList(todoListEntries...)
}
