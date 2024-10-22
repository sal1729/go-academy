package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func setActionAndFilter(c *CrudRequest, action string, task string, status string) error {
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

func getCrudRequestFromUserInput(input UserInput) (CrudRequest, error) {
	var crudRequest CrudRequest
	var action string
	var task string
	var err error

	status := input.Status
	if input.Create != "" {
		action = "create"
		task = input.Create
		err = setActionAndFilter(&crudRequest, action, task, status)
	}
	if input.List != "" {
		action = "list"
		task = input.List
		err = setActionAndFilter(&crudRequest, action, task, status)
	}
	if input.Update != "" {
		action = "update"
		task = input.Update
		err = setActionAndFilter(&crudRequest, action, task, status)
	}
	if input.Delete != "" {
		action = "delete"
		task = input.Delete
		err = setActionAndFilter(&crudRequest, action, task, input.Status)
	}

	if action == "" {
		err = errors.New("no action specified")
	}

	return crudRequest, err
}

func handleRequest(todoList []TodoItem, req CrudRequest) error {
	// TODO Handle unhappy path where there was nothing to update/delete/list
	switch req.Action {
	case "create":
		newEntry, createErr := createTask(todoList, req.Task, req.Status)
		if createErr != nil {
			return createErr
		}
		fmt.Println("Created entry:\n", newEntry)
	case "list":
		list, listErr := readFromList(todoList, req.Task, req.Status)
		if listErr != nil {
			return listErr
		}
		fmt.Println("Listing entries:", list)
	case "update":
		updates, updateErr := updateListItems(todoList, req.Task, req.Status)
		if updateErr != nil {
			return updateErr
		}
		fmt.Println("Updated entries:", updates)
	case "delete":
		deletions, deleteErr := deleteFromList(todoList, req.Task, req.Status)
		if deleteErr != nil {
			return deleteErr
		}
		fmt.Println("Deleted entries:", deletions)
	default:
		return fmt.Errorf("unknown action: %s", req.Action)
	}
	return nil
}

func createTask(todoList []TodoItem, task string, status string) ([]TodoItem, error) {
	fmt.Printf("Creating task %s, status %s\n", task, status)
	newEntry := TodoItem{}

	if taskExists(todoList, task) {
		return []TodoItem{newEntry}, fmt.Errorf("cannot create - task %s already exists", task)
	}
	newEntry.Task = task
	newEntry.Status = status

	return []TodoItem{newEntry}, nil
}

// readFromList does not error if no such task is found.
func readFromList(todoList []TodoItem, taskname string, status string) ([]TodoItem, error) {
	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Reading task %s, status %s\n", taskname, statusString)

	if taskname == "all" && status == "" {
		return todoList, nil
	}
	if taskname == "all" { // && status != "", implicitly
		return filterByStatus(todoList, status), nil
	}

	// Implicitly, if task != "all"
	task, _ := getTaskByName(todoList, taskname)
	return []TodoItem{task}, nil
}

func updateListItems(todoList []TodoItem, task string, status string) ([]TodoItem, error) {
	fmt.Printf("Updating task %s, status %s\n", task, status)
	var updates []TodoItem

	for i := range todoList {
		if task == "all" {
			currentStatus, newStatus, statusErr := decomposeStatusString(status)
			if statusErr != nil {
				return updates, statusErr
			}
			if todoList[i].Status == currentStatus {
				todoList[i].Status = newStatus
				updates = append(updates, todoList[i])
			}
		} else if todoList[i].Task == task {
			todoList[i].Status = status
			updates = append(updates, todoList[i])
			break // Here we assume task name is unique
		}
	}

	return updates, nil
}

func deleteFromList(todoList []TodoItem, task string, status string) ([]TodoItem, error) {
	taskString := "all"
	if task != "" {
		taskString = task
	}
	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Deleting task %s, status %s\n", taskString, statusString)
	var deletions []TodoItem

	// Use a single pass to filter entries in-place.
	for i := 0; i < len(todoList); {
		entry := todoList[i]
		// Conditions for deletion
		if (task == "all" && status == "") ||
			(task == "all" && status != "" && entry.Status == status) ||
			(task != "all" && entry.Task == task) {
			deletions = append(deletions, entry)
			// Remove the entry by swapping with the last element and slicing off the last item
			// Gotta confess, this was pure ChatGPT but it's pretty nifty
			todoList[i] = todoList[len(todoList)-1]
			todoList = todoList[:len(todoList)-1]
		} else {
			i++ // Move to the next item only if no deletion happened - if there was deletion there's a new entry to check in slot i
		}
	}

	return deletions, nil
}

func taskExists(entries []TodoItem, task string) bool {
	for _, entry := range entries {
		if entry.Task == task {
			return true
		}
	}
	return false
}

func getTaskByName(todoList []TodoItem, taskName string) (TodoItem, error) {
	for _, entry := range todoList {
		if entry.Task == taskName {
			return entry, nil // Returns first task with requested name, we're assuming this is unique
		}
	}
	return TodoItem{}, fmt.Errorf("task not found: %s", taskName)
}

func filterByStatus(todoList []TodoItem, status string) []TodoItem {
	var filteredList []TodoItem
	for _, entry := range todoList {
		if entry.Status == status {
			filteredList = append(filteredList, entry)
		}
	}
	return filteredList
}

func decomposeStatusString(input string) (string, string, error) {
	parts := strings.Split(input, "*TO*")

	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid input, expecting 'status*TO*newStatus'")
	}

	return parts[0], parts[1], nil
}
