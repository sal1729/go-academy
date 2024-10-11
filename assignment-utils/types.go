package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type NameType struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (name NameType) String() string {
	// Display the full-name : <full-name>, middle-name : <middle-name> and surname : <surname> on 3 separate lines.
	return fmt.Sprintf("full-name: %s %s %s\nmiddle-name: %s\nsurname: %s\n", name.FirstName, name.MiddleName, name.LastName, name.MiddleName, name.LastName)
}

type DobType struct {
	Year  int        `json:"year"`
	Month time.Month `json:"month"`
	Day   int        `json:"day"`
}

type DobString string

var dobPattern = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func NewDobString(input string) (DobString, error) {
	// Basic validation
	parts := strings.Split(input, "-")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format")
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format, error converting string to integer: %s", err)
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format, error converting string to integer: %s", err)
	}
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format, error converting string to integer: %s", err)
	}

	if month < 1 || month > 12 {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format, month must be between 1 and 12")
	}
	if day < 1 || day > 31 {
		return "", fmt.Errorf("invalid date of birth, expecting YYYY-MM-DD format, day must be between 1 and 31")
	}

	// Month/Day pair validation
	shortMonths := []int{2, 4, 6, 9, 11}
	includes := func(slice []int, x int) bool {
		for _, v := range slice {
			if v == x {
				return true
			}
		}
		return false
	}
	if includes(shortMonths, month) && day == 31 {
		return "", fmt.Errorf("invalid date of birth, the specified month has less than 31 days")
	}

	// February validation
	isLeapYear := func(year int) bool {
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			return true
		}
		return false
	}
	if month == 2 {
		if (isLeapYear(year) && day > 29) || day > 28 {
			return "", fmt.Errorf("invalid date of birth, February has less than %d days", day)
		}
	}

	// Final check
	dobString := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	if !dobPattern.MatchString(dobString) {
		return "", fmt.Errorf("invalid date of birth, inputs do not stringify correctly")
	}

	return DobString(dobString), nil
}

func (dob DobType) ToString() (DobString, error) {
	return NewDobString(fmt.Sprintf("%d-%d-%d", dob.Year, dob.Month, dob.Day))
}

func (dob DobString) ToDobType() DobType {
	parts := strings.Split(string(dob), "-")
	Year, _ := strconv.Atoi(parts[0])
	Month, _ := strconv.Atoi(parts[1])
	Day, _ := strconv.Atoi(parts[2])
	return DobType{Year, time.Month(Month), Day}
}

type StudentDetails struct {
	Name NameType  `json:"name"`
	Dob  DobString `json:"dob"`
}
