package main

import (
	"fmt"
	"go-academy/utils"
	"log"
)

func main() {
	// The programme will display two strings in the same stdout line
	// First, set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("nameRequest: ")
	log.SetFlags(0)

	requests := [3]string{
		"Please enter your name. First name:",
		"Please enter your name. Any middle names, hit enter to skip:",
		"Please enter your name. Last name:",
	}

	inputs, inputError := utils.RequestInputs(requests[0:])
	if inputError != nil {
		log.Fatal(inputError)
	}

	// If the middle name was empty, we remove it.
	if inputs[1] == "" {
		inputs = append(inputs[:1], inputs[2])
	}

	displayOptions := &utils.DisplayOptions{Delimiter: " ", LineEnd: ""}
	name, displayErr := utils.StringDisplay(inputs, displayOptions)
	if displayErr != nil {
		log.Fatal(displayErr)
	}
	fmt.Printf("Hello, %s\n", name)
}
