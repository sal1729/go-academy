package main

import (
	"bufio"
	"fmt"
	"go-academy/utils"
	"log"
	"os"
)

// Started as a straight copy of the program from assignment-02
func main() {
	// The programme will display two strings in the same stdout line
	// First, set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("nameRequest: ")
	log.SetFlags(0)

	requests := [3]string{
		"Please enter your name. First name: ",
		"Please enter your name. Any middle names, hit enter to skip: ",
		"Please enter your name. Last name: ",
	}

	reader := bufio.NewReader(os.Stdin)
	inputs, inputError := utils.RequestInputs(requests[0:], reader)
	if inputError != nil {
		log.Fatal(inputError)
	}

	// Create name struct from inputs
	name := utils.NameType{
		FirstName:  inputs[0],
		MiddleName: inputs[1],
		LastName:   inputs[2],
	}

	// Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines.
	fmt.Printf("full-name: %s %s %s\n", name.FirstName, name.MiddleName, name.LastName)
	fmt.Printf("middle-name: %s\n", name.MiddleName)
	fmt.Printf("surname: %s\n", name.LastName)
}
