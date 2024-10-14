package todo_app_functions

import (
	"encoding/json"
	"fmt"
)

func PrintList(list ...ListEntry) {
	for _, entry := range list {
		fmt.Println(entry)
	}
}

func ListAsJson(list ...ListEntry) (string, error) {
	todoList := TodoList{
		Entries: list,
	}
	jsonData, err := json.MarshalIndent(todoList, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
