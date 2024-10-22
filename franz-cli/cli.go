// This is exercise 17 again
package main

import (
	"flag"
	"log"
)

// Based on exercise-17.go
// Run `go build -o franz cli.go` to construct the executable
// Our app is called franz, after Franz Liszt.
// Example commands
// CREATE: `./franz -create "New Task Name" -status "In Progress"`
// READ: `./franz -list`, `./franz -list "Task Name"`, `./franz -list -status "Status"`
// UPDATE: `./franz -update "Task Name" -status "New Status"`
// DELETE: `./franz -delete "Task Name"`, `./franz -delete -status "Status"`
func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("franz: ")
	log.SetFlags(0)
	// Also set up a function for catching errors
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Option flags
	// filename := flag.String("filename", "todo-list.json", "filename to read")
	create := flag.String("create", "", "task to create, must pass with -status")
	status := flag.String("status", "", "status to create/read/update/delete, must be passed with other flags")
	list := flag.String("list", "", "pass with all to list all tasks, with taskname to list single task, with all and -status to filter on status")
	update := flag.String("update", "", "task to update, must pass with -status")
	del := flag.String("delete", "", "pass with all to delete all tasks, with taskname to delete single task, with all and -status to delete by status")
	flag.Parse()

	// Convert user input to CrudRequest
	input := UserInput{Create: *create, Status: *status, List: *list, Update: *update, Delete: *del}
	request, err := getCrudRequestFromUserInput(input)
	catchError(err)

	// Get data
	todoList := []TodoItem{
		{Task: "Feed the floor", Status: "To Do"},
		{Task: "Sweep the dishes", Status: "In Progress"},
		{Task: "Rock the rug", Status: "In Progress"},
		{Task: "Scrub the fishes", Status: "To Do"},
		{Task: "Vacuum the lawn", Status: "Done"},
		{Task: "Bathe the mat", Status: "Done"},
		{Task: "Mop the baby", Status: "In Progress"},
		{Task: "Mow the cat", Status: "Blocked"},
		{Task: "Stop! Look!", Status: "To Do"},
		{Task: "Buy the book", Status: "Blocked"},
	}

	// Handle request
	requestError := handleRequest(todoList, request)
	catchError(requestError)

	log.Println("Request handled successfully:", request)
}
