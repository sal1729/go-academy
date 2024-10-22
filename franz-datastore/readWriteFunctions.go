package franz_datastore

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// TODO Find a better solution to this
const (
	datastore = "/Users/sarah.liddell/bench/golang/go-academy/franz-datastore"
)

type SavedList struct {
	Entries []ListItem `json:"entries"`
}

func GetData() ([]ListItem, error) {
	var todoList SavedList

	dataFile := filepath.Join(datastore, fmt.Sprintf("data/%s", "todo-list.json"))

	todoListData, err := jsonBytesFromFile(dataFile)
	if err != nil {
		return todoList.Entries, err
	}

	err = json.Unmarshal(todoListData, &todoList)
	if err != nil {
		return todoList.Entries, err
	}

	return todoList.Entries, nil
}

// This is a copy of JsonBytesFromFile from JsonTools.go
func jsonBytesFromFile(filename string) ([]byte, error) {
	var byteData []byte
	file, err := os.Open(filename)
	if err != nil {
		return byteData, err
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file %s: %v", filename, closeErr)
		}
	}()

	byteData, err = io.ReadAll(file)
	if err != nil {
		return byteData, err
	}

	return byteData, nil
}

func SaveData(todoList *[]ListItem) error {
	dataFile := filepath.Join(datastore, fmt.Sprintf("data/%s", "todo-list.json"))

	byteCount, err := listToFile(dataFile, *todoList...)
	if err != nil {
		return err
	}
	fmt.Printf("Saved %s, %d bytes\n", dataFile, byteCount)
	return nil
}

// This is adapted from functions in ToDoList.go
func listToFile(filename string, list ...ListItem) (int, error) {
	jsonBytes, err := json.MarshalIndent(SavedList{Entries: list}, "", "  ")
	if err != nil {
		return 0, err
	}

	file, err := os.Create(filename)
	if err != nil {
		return 0, err
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close file %s: %v\n", filename, closeErr)
		}
	}()

	b, err := file.Write(jsonBytes)
	if err != nil {
		return 0, err
	}
	return b, nil
}
