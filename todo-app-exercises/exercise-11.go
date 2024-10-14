package main

import (
	"log"
	todo "todo_app_functions"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("todo_app: ")
	log.SetFlags(0)
	// Also set up an function for catching errors
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

	jsonString, err := todo.ListAsJson(tasks...)
	catchError(err)
	log.Println(jsonString)
}
