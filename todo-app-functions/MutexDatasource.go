package todo_app_functions

// I took a copy of the original Datasource.go so that I didn't fry any of the earlier exercises
// when trying to make things nice and concurrent

import (
	"encoding/json"
	"fmt"
	"sync"
)

type MutexDatasource interface {
	Create(task, status string) (ListEntry, error)
	Read(task, status string) (Entries, error)
	Update(task, status string) (Entries, error)
	Delete(task, status string) (Entries, error)
}

type LocalMuDatasource struct {
	Filepath string
	mu       sync.RWMutex
}

func (d *LocalMuDatasource) getData() (TodoList, error) {
	var todoList TodoList
	todoListData, readErr := JsonBytesFromFile(d.Filepath)
	if readErr != nil {
		return todoList, readErr
	}

	jsonErr := json.Unmarshal(todoListData, &todoList)
	if jsonErr != nil {
		return todoList, readErr
	}

	return todoList, nil
}

func (d *LocalMuDatasource) saveData(todoList TodoList) error {
	byteCount, err := ListToFile(d.Filepath, todoList.Entries...)
	if err != nil {
		return err
	}
	fmt.Printf("Saved %s, %d bytes\n", d.Filepath, byteCount)
	return nil
}

func (d *LocalMuDatasource) Create(task, status string) (ListEntry, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Printf("Creating task %s, status %s\n", task, status)
	newEntry := ListEntry{}
	todoList, dataErr := d.getData()
	if dataErr != nil {
		return newEntry, dataErr
	}

	if taskExists(todoList.Entries, task) {
		return newEntry, fmt.Errorf("cannot create - task %s already exists", task)
	}
	newEntry = ListEntry{
		Task:   task,
		Status: status,
	}
	todoList.Entries = append(todoList.Entries, newEntry)

	saveErr := d.saveData(todoList)
	if saveErr != nil {
		return newEntry, saveErr
	}
	return newEntry, nil
}

func (d *LocalMuDatasource) Read(task, status string) (Entries, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Reading task %s, status %s\n", task, statusString)
	var list []ListEntry
	todoList, dataErr := d.getData()
	if dataErr != nil {
		return list, dataErr
	}

	if task == "all" && status == "" {
		return todoList.Entries, nil
	}
	if task == "all" { // && status != "", implicitly
		return FilterByStatus(todoList.Entries, status), nil
	}

	taskEntry, _ := GetTaskByName(todoList.Entries, task)
	list = append(list, taskEntry)
	return list, nil
}

func (d *LocalMuDatasource) Update(task, status string) (Entries, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Printf("Updating task %s, status %s\n", task, status)
	var updates []ListEntry
	todoList, dataErr := d.getData()
	if dataErr != nil {
		return updates, dataErr
	}

	for i := range todoList.Entries {
		if task == "all" {
			currentStatus, newStatus, statusErr := decomposeStatusString(status)
			if statusErr != nil {
				return updates, statusErr
			}
			if todoList.Entries[i].Status == currentStatus {
				todoList.Entries[i].Status = newStatus
				updates = append(updates, todoList.Entries[i])
			}
		} else if todoList.Entries[i].Task == task {
			todoList.Entries[i].Status = status
			updates = append(updates, todoList.Entries[i])
			break // Here we assume task name is unique
		}
	}

	saveErr := d.saveData(todoList)
	if saveErr != nil {
		return updates, saveErr
	}
	return updates, nil
}

func (d *LocalMuDatasource) Delete(task, status string) (Entries, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	taskString := "all"
	if task != "" {
		taskString = task
	}
	statusString := "all"
	if status != "" {
		statusString = status
	}
	fmt.Printf("Deleting task %s, status %s\n", taskString, statusString)
	var deletions []ListEntry
	todoList, dataErr := d.getData()
	if dataErr != nil {
		return deletions, dataErr
	}

	// Use a single pass to filter entries in-place.
	for i := 0; i < len(todoList.Entries); {
		entry := todoList.Entries[i]
		// Conditions for deletion
		if (task == "all" && status == "") ||
			(task == "all" && status != "" && entry.Status == status) ||
			(task != "all" && entry.Task == task) {
			deletions = append(deletions, entry)
			// Remove the entry by swapping with the last element and slicing off the last item
			// Gotta confess, this was pure ChatGPT but it's pretty nifty
			todoList.Entries[i] = todoList.Entries[len(todoList.Entries)-1]
			todoList.Entries = todoList.Entries[:len(todoList.Entries)-1]
		} else {
			i++ // Move to the next item only if no deletion happened - if there was deletion there's a new entry to check in slot i
		}
	}

	saveErr := d.saveData(todoList)
	if saveErr != nil {
		return deletions, saveErr
	}
	return deletions, nil
}
