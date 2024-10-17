package main

import (
	"log"
	"net/http"
)

// Based on https://go.dev/doc/articles/wiki/
// Run `go build -o franz_web exercise-18.go` to construct the executable
// Our app is called franz, after Franz Liszt.
func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
