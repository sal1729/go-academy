package main

import (
	"bufio"
	"fmt"
	bbAge "github.com/bearbin/go-age"
	"go-academy/utils"
	"os"
	"strconv"
	"time"
)

// TODO I must be missing something here because the assignment description implied I should have done more work
// Create a program that calculates the age of a person given their date of birth.
// [Variables][Methods][Arrays][Slices][For Loops][Package Usage]
func main() {
	requests := []string{
		"What is your date of birth? Year: ",
		"What is your date of birth? Month: ",
		"What is your date of birth? Day: ",
	}
	reader := bufio.NewReader(os.Stdin)
	inputs, _ := utils.RequestInputs(requests, reader)

	// TODO Again, zero error handling here. Bad stuff happens if we pass illegal months/days
	// I did a bunch of dob error handling later though...
	year, _ := strconv.Atoi(inputs[0])
	month, _ := strconv.Atoi(inputs[1])
	day, _ := strconv.Atoi(inputs[2])

	dob := utils.DobType{
		Year:  year,
		Month: time.Month(month),
		Day:   day,
	}

	birthday := time.Date(dob.Year, dob.Month, dob.Day, 0, 0, 0, 0, time.UTC)
	age := bbAge.Age(birthday) //
	fmt.Println("Age: ", age)
}
