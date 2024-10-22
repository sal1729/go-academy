package franz_datastore

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidateCrudRequest might be a bit out of data now
func ValidateCrudRequest(req CrudRequest) error {
	action := req.Action
	task := req.Task
	status := req.Status

	if task == "" {
		return fmt.Errorf("task cannot be empty")
	}
	if task == "all" && action == "create" {
		return fmt.Errorf("cannot create 'all' - reserved keyword")
	}

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
	return nil
}

// CreateTask errors if task already exists
func CreateTask(todoList *[]ListItem, req CrudRequest) ([]ListItem, error) {
	err := ValidateCrudRequest(req)
	if err != nil {
		return []ListItem{}, err
	}

	task := req.Task
	status := req.Status
	fmt.Printf("Creating task %s, status %s\n", task, status)
	newEntry := ListItem{}

	if taskExists(*todoList, task) {
		return []ListItem{newEntry}, fmt.Errorf("cannot create - task %s already exists", task)
	}

	// Create task, modify data in memory, return
	newEntry.Task = task
	newEntry.Status = status
	*todoList = append(*todoList, newEntry)
	return []ListItem{newEntry}, nil
}

// ReadFromList does not error if no such task is found.
func ReadFromList(todoList *[]ListItem, req CrudRequest) ([]ListItem, error) {
	err := ValidateCrudRequest(req)
	if err != nil {
		return []ListItem{}, err
	}

	taskname := req.Task
	status := req.Status

	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Reading task %s, status %s\n", taskname, statusString)

	if taskname == "all" && status == "" {
		return *todoList, nil
	}
	if taskname == "all" { // && status != "", implicitly
		return filterByStatus(*todoList, status), nil
	}

	// Implicitly, if task != "all"
	task, _ := getTaskByName(*todoList, taskname)
	return []ListItem{task}, nil
}

// UpdateListItems errors if there is something wrong with the request
// but it doesn't error if it can't find a task to update
func UpdateListItems(todoList *[]ListItem, req CrudRequest) ([]ListItem, error) {
	err := ValidateCrudRequest(req)
	if err != nil {
		return []ListItem{}, err
	}

	task := req.Task
	status := req.Status
	fmt.Printf("Updating task %s, status %s\n", task, status)
	var updates []ListItem

	for i := range *todoList {
		if task == "all" {
			currentStatus, newStatus, statusErr := decomposeStatusString(status)
			if statusErr != nil {
				return updates, statusErr
			}
			if (*todoList)[i].Status == currentStatus {
				(*todoList)[i].Status = newStatus
				updates = append(updates, (*todoList)[i])
			}
		} else if (*todoList)[i].Task == task {
			(*todoList)[i].Status = status
			updates = append(updates, (*todoList)[i])
			break // Here we assume task name is unique
		}
	}

	return updates, nil
}

// DeleteFromList does not throw errors, even if the requested task does not exist
func DeleteFromList(todoList *[]ListItem, req CrudRequest) ([]ListItem, error) {
	err := ValidateCrudRequest(req)
	if err != nil {
		return []ListItem{}, err
	}

	task := req.Task
	status := req.Status

	taskString := "all"
	if task != "" {
		taskString = task
	}
	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Deleting task %s, status %s\n", taskString, statusString)
	var deletions []ListItem

	// Use a single pass to filter entries in-place.
	for i := 0; i < len(*todoList); {
		entry := (*todoList)[i]
		// Conditions for deletion
		if (task == "all" && status == "") ||
			(task == "all" && status != "" && entry.Status == status) ||
			(task != "all" && entry.Task == task) {
			deletions = append(deletions, entry)
			// Remove the entry by swapping with the last element and slicing off the last item
			// Gotta confess, this was pure ChatGPT but it's pretty nifty
			(*todoList)[i] = (*todoList)[len(*todoList)-1]
			*todoList = (*todoList)[:len(*todoList)-1]
		} else {
			i++ // Move to the next item only if no deletion happened - if there was deletion there's a new entry to check in slot i
		}
	}

	return deletions, nil
}

func taskExists(entries []ListItem, task string) bool {
	for _, entry := range entries {
		if entry.Task == task {
			return true
		}
	}
	return false
}

func getTaskByName(todoList []ListItem, taskName string) (ListItem, error) {
	for _, entry := range todoList {
		if entry.Task == taskName {
			return entry, nil // Returns first task with requested name, we're assuming this is unique
		}
	}
	return ListItem{}, fmt.Errorf("task not found: %s", taskName)
}

func filterByStatus(todoList []ListItem, status string) []ListItem {
	var filteredList []ListItem
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
