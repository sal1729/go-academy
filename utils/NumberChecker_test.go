package utils

import (
	"testing"
)

// TestNumberCheckerNumericOnly calls NumberChecker with a bad string,
// checking for the expected error.
func TestNumberCheckerNumericOnly(t *testing.T) {
	nonIntegerInput := "Gladys"
	response, err := NumberChecker(nonIntegerInput, 0, 10)
	if response != "" || err == nil {
		t.Fatalf(`NumberChecker(nonIntegerInput,0,10) = %q, %v, want "", error`, response, err)
	}
}

// TestNumberCheckerLessThan returns informative string, checking
//
//	for a valid return value.
func TestNumberCheckerLessThan(t *testing.T) {
	integerInput := "-1"
	response, err := NumberChecker(integerInput, 0, 10)
	expected := "Your number is less than 0"
	if response != expected || err != nil {
		t.Fatalf(`NumberChecker("-1") = %q, %v, want match for %#q, nil`, response, err, expected)
	}
}

// TestNumberCheckerGreaterThan returns informative string, checking
//
//	for a valid return value.
func TestNumberCheckerGreaterThan(t *testing.T) {
	integerInput := "11"
	response, err := NumberChecker(integerInput, 0, 10)
	expected := "Your number is greater than 10"
	if response != expected || err != nil {
		t.Fatalf(`NumberChecker("11") = %q, %v, want match for %#q, nil`, response, err, expected)
	}
}

// TestNumberCheckerValid returns informative string, checking
//
//	for a valid return value.
func TestNumberCheckerValid(t *testing.T) {
	integerInput := "1"
	response, err := NumberChecker(integerInput, 0, 10)
	expected := "Your number is valid"
	if response != expected || err != nil {
		t.Fatalf(`NumberChecker("1") = %q, %v, want match for %#q, nil`, response, err, expected)
	}
}
