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
	log.SetPrefix("display: ")
	log.SetFlags(0)

	// The inputs
	inputs := []string{"This is a demo sentence", "this is an additional clause"}

	// Get the formatted string
	displayOptions := &utils.DisplayOptions{Delimiter: " - ", LineEnd: "!"}
	displayString, err := utils.StringDisplay(inputs, displayOptions)
	//displayString, err := utils.StringDisplay(inputs, nil)
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print string
	fmt.Println(displayString)
}
