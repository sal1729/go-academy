package todo_app_functions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// JsonBytesFromFile is a more generic version of getStudentDetails from assignment-utils/Register.go
// It returns the file contents in a byte array
func JsonBytesFromFile(filename string) ([]byte, error) {
	var zero []byte
	file, openErr := os.Open(filename)
	if openErr != nil {
		return zero, openErr
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file %s: %v", filename, closeErr)
		}
	}()

	byteData, readErr := io.ReadAll(file)
	if readErr != nil {
		return zero, readErr
	}

	return byteData, nil
}

// PrettyPrintJson will pretty print any old json
func PrettyPrintJson(jsonData []byte) error {
	var prettyJson bytes.Buffer

	err := json.Indent(&prettyJson, jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to pretty-print JSON: %v", err)
	}

	fmt.Println(prettyJson.String())
	return nil
}

// StreamEntriesFromFile allows us to stream the first n tasks from a json file
// containing a TodoList (as json)
// TODO test
func StreamEntriesFromFile(filename string, n int) ([]ListEntry, error) {
	file, openErr := os.Open(filename)
	if openErr != nil {
		return nil, fmt.Errorf("error opening file: %w", openErr)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file %s: %v", filename, closeErr)
		}
	}()

	decoder := json.NewDecoder(file)
	_, braceErr := decoder.Token()
	if braceErr != nil {
		return nil, fmt.Errorf("error reading opening brace: %w", braceErr)
	}

	for decoder.More() {
		token, tokenErr := decoder.Token()
		if tokenErr != nil {
			return nil, fmt.Errorf("error reading token: %w", tokenErr)
		}

		if key, ok := token.(string); ok && key == "entries" {
			break
		}
	}

	_, arrayErr := decoder.Token()
	if arrayErr != nil {
		return nil, fmt.Errorf("error reading array: %w", arrayErr)
	}

	var entries []ListEntry
	count := 0

	for decoder.More() && count < n {
		var entry ListEntry
		if entryErr := decoder.Decode(&entry); entryErr != nil {
			return nil, fmt.Errorf("error decoding entry: %w", entryErr)
		}
		entries = append(entries, entry)
		count++
	}

	return entries, nil
}
