package todo_app_functions

import "fmt"

type ListEntry struct {
	Task   string `json:"task"`
	Status string `json:"status"` // Adding this has probably fired some of the early exercises
}

func (e ListEntry) String() string {
	if e.Status != "" {
		return fmt.Sprintf("%s - %s", e.Task, e.Status)
	}
	return e.Task
}

type TodoList struct {
	Entries []ListEntry `json:"entries"`
}

type JsonOutput interface {
	~string | ~[]byte
}
