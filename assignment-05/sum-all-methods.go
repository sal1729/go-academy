package main

import (
	"bufio"
	"fmt"
	"go-academy/utils"
	"log"
	"os"
)

func main() {
	// The programme will display two strings in the same stdout line
	// First, set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("nameRequest: ")
	log.SetFlags(0)

	// First get three sets of numbers
	reader := bufio.NewReader(os.Stdin)
	n1, requestErr1 := utils.RequestNumbers(3, 0, 9, reader)
	if requestErr1 != nil {
		log.Fatal(requestErr1)
	}
	n2, requestErr2 := utils.RequestNumbers(3, 10, 99, reader)
	if requestErr2 != nil {
		log.Fatal(requestErr2)
	}
	n3, requestErr3 := utils.RequestNumbers(3, 100, 999, reader)
	if requestErr3 != nil {
		log.Fatal(requestErr3)
	}

	fmt.Println("Input groups: ", n1, n2, n3)

	// Create typed arrays
	// By not catching error here we permit silent failure 😬
	units, _ := utils.NewUnits(n1)
	tens, _ := utils.NewTens(n2)
	hundreds, _ := utils.NewHundreds(n3)

	// Use methods to sum
	uSum := units.Sum()
	tSum := tens.Sum()
	hSum := hundreds.Sum()

	fmt.Println("Group sums: ", uSum, tSum, hSum)
	fmt.Println("Total sum:", uSum+tSum+hSum)

}
