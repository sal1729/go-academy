package main

import (
	"regexp"
	"testing"
)

// TestStringDisplayDefaults calls StringDisplay without displayOptions, checking
// for a valid return value.
func TestStringDisplayDefaults(t *testing.T) {
	inputOne := "inputOne"
	inputTwo := "inputTwo"
	want := regexp.MustCompile(`^` + inputOne + ", " + inputTwo + "." + `$`)
	msg, err := StringDisplay(inputOne, inputTwo, nil)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("inputOne","inputTwo",nil) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestStringDisplayOptions calls StringDisplay with displayOptions, checking
// for a valid return value.
func TestStringDisplayOptions(t *testing.T) {
	inputOne := "inputOne"
	inputTwo := "inputTwo"
	displayOptions := &DisplayOptions{delimiter: " - ", lineEnd: "!"}
	want := regexp.MustCompile(`^` + inputOne + " - " + inputTwo + "!" + `$`)
	msg, err := StringDisplay(inputOne, inputTwo, displayOptions)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("inputOne","inputTwo",nil) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestStringDisplayEmptyOne calls StringDisplay with an empty string in the first input,
// checking for an error.
func TestHelloEmptyOne(t *testing.T) {
	msg, err := StringDisplay("", "inputTwo", nil)
	if msg != "" || err == nil {
		t.Fatalf(`StringDisplay("","inputTwo",nil) = %q, %v, want "", error`, msg, err)
	}
}

// TestStringDisplayEmptyTwo calls StringDisplay with an empty string in the second input,
// checking for an error.
func TestHelloEmptyTwo(t *testing.T) {
	msg, err := StringDisplay("inputOne", "", nil)
	if msg != "" || err == nil {
		t.Fatalf(`StringDisplay("inputOne","",nil) = %q, %v, want "", error`, msg, err)
	}
}

// TestStringDisplayEmptyBoth calls StringDisplay with an empty string in the both inputs,
// checking for an error.
func TestHelloEmptyBoth(t *testing.T) {
	msg, err := StringDisplay("", "", nil)
	if msg != "" || err == nil {
		t.Fatalf(`StringDisplay("","",nil) = %q, %v, want "", error`, msg, err)
	}
}
