package main

import (
	"go-academy/utils"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	// TODO Use log more consistently across other programs
	log.SetPrefix("city-display: ")
	log.SetFlags(0)
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	cities := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}
	byteSizedCities := []byte(strings.Join(cities, "\n"))

	// Create temporary file and write to it
	tempFile, err := os.CreateTemp("", "cities")
	catchError(err)
	log.Println("Temporary file:", tempFile.Name())

	fileSize, err := tempFile.Write(byteSizedCities)
	catchError(err)
	log.Printf("Wrote %d bytes to temporary file %s.\n", fileSize, tempFile.Name())

	// Close so we can read in the next step
	err = tempFile.Close()
	catchError(err)

	// Read cities from temporary file
	fileBytes, err := os.ReadFile(tempFile.Name())
	catchError(err)
	fileLines := strings.Split(string(fileBytes), "\n")

	// Display cities in alphabetical order
	sort.Slice(fileLines, func(i, j int) bool {
		return fileLines[i] < fileLines[j]
	})
	displayOpts := utils.DisplayOptions{
		Delimiter: ", ",
		LineEnd:   "",
	}
	cityString, err := utils.StringDisplay(fileLines, &displayOpts)
	catchError(err)
	log.Println("Cities:", cityString)

	// Tidy up - this should happen automatically, but let's be explicit #learning
	defer func() {
		if err := os.Remove(tempFile.Name()); err != nil {
			log.Printf("Warning: failed to remove temp file %s: %v", tempFile.Name(), err)
		} else {
			log.Printf("Removed temporary file %s.\n", tempFile.Name())
		}
	}()
}
