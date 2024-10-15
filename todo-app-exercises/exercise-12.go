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

	tasks := []todo.ListEntry{
		{Task: "Feed the floor"},
		{Task: "Sweep the dishes"},
		{Task: "Rock the rug"},
		{Task: "Scrub the fishes"},
		{Task: "Vacuum the lawn"},
		{Task: "Bathe the mat"},
		{Task: "Mop the baby"},
		{Task: "Mow the cat"},
		{Task: "Stop! Look!"},
		{Task: "Buy the book"},
	}

	dir, err := os.Getwd()
	catchError(err)

	filename := filepath.Join(dir, "files/todo-list.json")
	byteCount, err := todo.ListToFile(filename, tasks...)
	catchError(err)
	log.Printf("Saved %s, %d bytes", filename, byteCount)
}
