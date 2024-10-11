package utils

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

// TestRequestInputs request an input and returns an array of trimmed strings
// for a valid return value.
func TestRequestInputs(t *testing.T) {
	requests := []string{"Test input request one:", "Test input request two:"}
	testInputs := "INPUT ONE\nINPUT TWO\n"
	testReader := bufio.NewReader(strings.NewReader(testInputs))
	expected := []string{"INPUT ONE", "INPUT TWO"}

	results, _ := RequestInputs(requests, testReader)
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("TestRequestInputs failed, expected %v, got %v", expected, results)
	}
}
