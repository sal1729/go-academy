package main

import (
	franz "franz_datastore"
	"log"
	"net/http"
)

func main() {
	// Fire up request channel and datastore
	requests := make(chan franz.AccessRequest)
	go franz.Datastore(requests)

	// Start the HTTP server
	http.HandleFunc("/api", makeHandler(franzApi, requests))
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
