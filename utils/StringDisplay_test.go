package utils

import (
	"regexp"
	"testing"
)

// TestStringDisplayDefaults calls StringDisplay without displayOptions, checking
// for a valid return value.
func TestStringDisplayDefaults(t *testing.T) {
	inputs := []string{"inputOne", "inputTwo"}
	want := regexp.MustCompile(`^` + "inputOne, inputTwo" + "." + `$`)
	msg, err := StringDisplay(inputs, nil)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`StringDisplay(inputs ,nil) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestStringDisplayOptions calls StringDisplay with displayOptions, checking
// for a valid return value.
func TestStringDisplayOptions(t *testing.T) {
	inputs := []string{"inputOne", "inputTwo"}
	displayOptions := &DisplayOptions{Delimiter: " - ", LineEnd: "!"}
	want := regexp.MustCompile(`^` + "inputOne - inputTwo" + "!" + `$`)
	msg, err := StringDisplay(inputs, displayOptions)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`StringDisplay([],{Delimiter: " - ", LineEnd: "!"}) = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestStringDisplayEmptyInput calls StringDisplay with an empty input array,
// checking for an error.
func TestStringDisplayEmptyInput(t *testing.T) {
	var inputs []string
	msg, err := StringDisplay(inputs, nil)
	if msg != "" || err == nil {
		t.Fatalf(`StringDisplay([],nil) = %q, %v, want "", error`, msg, err)
	}
}
