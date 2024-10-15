package todo_app_functions

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func PrintList(list ...ListEntry) {
	for _, entry := range list {
		fmt.Println(entry)
	}
}

// ListAsJson 💯 The best function I've written so far
// TODO Make this fast-but-boring by only returning bytes 😢
func ListAsJson[T JsonOutput](list ...ListEntry) (T, error) {
	todoList := TodoList{
		Entries: list,
	}
	var zero T // Returning "zero" is a nice enough alternative to returning undefined 👌

	jsonData, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return zero, err
	}

	switch any(zero).(type) { // Yeah man, we can switch on the generic type 👌
	case string:
		return T(string(jsonData)), nil
	case []byte:
		return T(jsonData), nil
	default:
		return zero, fmt.Errorf("unsupported output type: %s", reflect.TypeOf(zero).String())
	}
}

func ListToFile(filename string, list ...ListEntry) (int, error) {
	jsonBytes, listErr := ListAsJson[[]byte](list...)
	var zero int
	if listErr != nil {
		return zero, listErr
	}

	file, createErr := os.Create(filename)
	if createErr != nil {
		return zero, createErr
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file %s: %v\n", filename, closeErr)
		} else {
			fmt.Printf("Closed file %s.\n", filename)
		}
	}()

	b, writeErr := file.Write(jsonBytes)
	if writeErr != nil {
		return zero, writeErr
	}
	return b, nil
}
