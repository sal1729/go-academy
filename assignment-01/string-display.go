package main

import (
	"fmt"
	"log"
)

type DisplayOptions struct {
	delimiter string
	lineEnd   string
}

func StringDisplay(s string, t string, opts *DisplayOptions) (string, error) {
	if s == "" || t == "" {
		return "", fmt.Errorf("expected two string inputs - input one: %s, input two: %s", s, t)
	}

	// Set display options, defaulting if not provided
	defaultDisplayOpts := DisplayOptions{delimiter: ", ", lineEnd: "."}
	delimiter := defaultDisplayOpts.delimiter
	lineEnd := defaultDisplayOpts.lineEnd
	if opts != nil && opts.delimiter != "" {
		delimiter = opts.delimiter
	}
	if opts != nil && opts.lineEnd != "" {
		lineEnd = opts.lineEnd
	}

	return fmt.Sprintf("%s%s%s%s", s, delimiter, t, lineEnd), nil
}

func main() {
	// The programme will display two strings in the same stdout line
	// First, set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("display: ")
	log.SetFlags(0)

	// The inputs
	inputOne := "This is a demo sentence"
	inputTwo := "this is an additional clause"

	// Get the formatted string
	displayOptions := &DisplayOptions{delimiter: " - ", lineEnd: "!"}
	displayString, err := StringDisplay(inputOne, inputTwo, displayOptions)
	//displayString, err := StringDisplay(inputOne, inputTwo, nil)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print string
	fmt.Println(displayString)
}
