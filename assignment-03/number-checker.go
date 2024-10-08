package main

import (
	"bufio"
	"fmt"
	"go-academy/utils"
	"log"
	"os"
)

func main() {
	// The programme will display information about the input number
	// First, set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("numberChecker: ")
	log.SetFlags(0)

	requests := [1]string{
		"Please enter a number: ",
	}
	reader := bufio.NewReader(os.Stdin)

	input, inputErr := utils.RequestInputs(requests[0:], reader)
	if inputErr != nil {
		log.Fatal(inputErr)
	}

	minVal := 1
	maxVal := 10
	numCheckResponse, numCheckErr := utils.NumberChecker(input[0], minVal, maxVal)
	if numCheckErr != nil {
		log.Fatal(numCheckErr)
	}
	fmt.Println(numCheckResponse)
}
