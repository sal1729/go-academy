package todo_app_functions

import (
	"reflect"
	"testing"
)

func TestListAsJsonString(t *testing.T) {
	input := []ListEntry{{Task: "Task 1"}, {Task: "Task 2"}}
	expectedOutput := "{\n  \"entries\": [\n    {\n      \"task\": \"Task 1\"\n    },\n    {\n      \"task\": \"Task 2\"\n    }\n  ]\n}"

	result, err := ListAsJson[string](input...)
	if err != nil {
		t.Errorf("did not expect an error, but got: %v", err)
		return
	}

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, got: %v", expectedOutput, result)
	}
}

func TestListAsJsonBytes(t *testing.T) {
	input := []ListEntry{{Task: "Task 1"}, {Task: "Task 2"}}
	expectedOutput := []byte("{\n  \"entries\": [\n    {\n      \"task\": \"Task 1\"\n    },\n    {\n      \"task\": \"Task 2\"\n    }\n  ]\n}")

	result, err := ListAsJson[[]byte](input...)
	if err != nil {
		t.Errorf("did not expect an error, but got: %v", err)
		return
	}

	if !reflect.DeepEqual(result, expectedOutput) {
		t.Errorf("Expected %v, got: %v", expectedOutput, result)
	}
}
