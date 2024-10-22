package main

import (
	franz "franz_datastore"
	"log"
	"net/http"
)

func main() {
	// Get data
	todoList := []franz.ListItem{
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

	// Open up channels

	// Start the HTTP server
	http.HandleFunc("/api", makeHandler(franzApi, &todoList))
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
