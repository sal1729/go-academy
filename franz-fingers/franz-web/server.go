package main

import (
	franz "franz_brain"
	"log"
	"net/http"
)

// This is adapted from franz-api/server.go and todo-app-web
func main() {
	// Fire up request channel and datastore
	requests := make(chan franz.AccessRequest)
	go franz.Datastore(requests)

	// Start the HTTP server, serving the api and the app webpage
	http.HandleFunc("/api", franz.MakeHandler(franz.Api, requests))
	http.HandleFunc("/franz", viewHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
