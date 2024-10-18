package todo_app_functions

import (
	"fmt"
	"regexp"
	"strings"
)

type ListEntry struct {
	Task   string `json:"task"`
	Status string `json:"status"` // Adding this has probably fired some of the early exercises
}

func (e ListEntry) String() string {
	return fmt.Sprintf("Task: %s Status: %s", e.Task, e.Status)
}

// This type was overkill tbh TODO ditch it
type TodoList struct {
	Entries []ListEntry `json:"entries"`
}

type JsonOutput interface {
	~string | ~[]byte
}

type Entries []ListEntry

func (e Entries) String() string {
	var result []string
	for _, todo := range e {
		result = append(result, fmt.Sprintf("Task: %s, Status: %s", todo.Task, todo.Status))
	}
	return "\n" + strings.Join(result, "\n")
}

type UserInput struct {
	Create string
	Status string
	List   string
	Update string
	Delete string
}

type CrudRequest struct {
	Action string `json:"action"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

func (c *CrudRequest) SetActionAndFilter(action, task, status string) error {
	if c.Action != "" {
		return fmt.Errorf("action (%s) is already set", c.Action)
	}
	c.Action = action

	if task == "" {
		return fmt.Errorf("task cannot be empty")
	}
	if task == "all" && action == "create" {
		return fmt.Errorf("cannot create 'all' - reserved keyword")
	}
	c.Task = task

	if status == "" && (action == "create" || action == "update") {
		return fmt.Errorf("cannot create or update without status")
	}
	if status != "" && (action == "list" || action == "delete") && task != "all" {
		return fmt.Errorf("cannot filter individual task by status when using list or delete")
	}
	re := regexp.MustCompile(`.*TO.*`)
	if !re.MatchString(status) && action == "update" && task == "all" {
		return fmt.Errorf("update status must be of the form 'status*TO*newStatus")
	}
	c.Status = status

	return nil
}
