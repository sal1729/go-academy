package main

import (
	"log"
	"os"
	"path/filepath"
	todo "todo_app_functions"
)

// exercise-13 would be better/improved/different/interesting-to-learn-from if it didn't read the
// whole file in to memory and instead streamed the tasks, stopping once we've hit the tenth
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
	tasksToRead := 10
	dir, dirErr := os.Getwd()
	catchError(dirErr)

	filename := filepath.Join(dir, "data/todo-list.json")
	todoListEntries, err := todo.StreamEntriesFromFile(filename, tasksToRead)
	catchError(err)

	log.Printf("The first %d tasks are", len(todoListEntries))
	todo.PrintList(todoListEntries...)
}
