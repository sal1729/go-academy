package todo_app_functions

import (
	"encoding/json"
	"fmt"
)

// Here we refactor the existing datasource to use goroutines and channels so that we don't have to
// keep opening and closing the file with the persisted data in.

type DataRequest struct {
	CrudRequest   CrudRequest
	ActionSuccess chan Result
	ListToView    chan []ListEntry
}

type Result struct {
	Success bool
	Error   error
}

type ConcurrentDatasource struct {
	Filepath string
	Requests chan DataRequest
}

func InitialiseDatasource(filepath string) *ConcurrentDatasource {
	d := &ConcurrentDatasource{
		Filepath: filepath,
		Requests: make(chan DataRequest),
	}

	// Start the datasource up - this will load the persisted data in to memory
	// and handle requests
	go d.run() // TODO Error handling here

	return d
}

func (d *ConcurrentDatasource) run() error {
	var entries []ListEntry

	// Load the initial data
	// TODO: Error handling
	entries, err := d.getData()
	if err != nil {
		return fmt.Errorf("unable to load data from %s:%v", d.Filepath, err)
	}

	for req := range d.Requests {
		action, task, status := req.CrudRequest.Action, req.CrudRequest.Task, req.CrudRequest.Status
		switch action {
		case "create":
			fmt.Printf("Creating task %s, status %s\n", task, status)
			newEntry := ListEntry{}

			if taskExists(entries, task) {
				err := fmt.Errorf("cannot create - task %s already exists", task)
				req.ActionSuccess <- Result{Success: false, Error: err}
				break
			}

			newEntry = ListEntry{
				Task:   task,
				Status: status,
			}

			entries = append(entries, newEntry)
			err := d.saveData(entries)
			if err != nil {
				req.ActionSuccess <- Result{Success: false, Error: err}
			}
			req.ActionSuccess <- Result{Success: true, Error: nil}

		case "read":
			statusString := "all"
			if status != "" {
				statusString = status
			}
			fmt.Printf("Reading task %s, status %s\n", task, statusString)

			if task == "all" && status == "" {
				req.ActionSuccess <- Result{Success: true, Error: nil}
				req.ListToView <- entries
				break
			}
			if task == "all" { // && status != "", implicitly
				filteredEntries := FilterByStatus(entries, status)
				req.ActionSuccess <- Result{Success: true, Error: nil}
				req.ListToView <- filteredEntries
				break
			}
			// if task != "all", implicitly
			taskEntry, _ := GetTaskByName(entries, task)
			req.ActionSuccess <- Result{Success: true, Error: nil}
			req.ListToView <- []ListEntry{taskEntry}

		case "update":
			fmt.Printf("Updating task %s, status %s\n", task, status)

		loop:
			for i := range entries {
				if task == "all" {
					currentStatus, newStatus, statusErr := decomposeStatusString(status)
					if statusErr != nil {
						req.ActionSuccess <- Result{Success: false, Error: statusErr}
						break loop // The internet assure me this should break out of both the loop and the case, but I am scepticals
					}
					if entries[i].Status == currentStatus {
						entries[i].Status = newStatus
					}
				} else if entries[i].Task == task {
					entries[i].Status = status
					break // Here we assume task name is unique
					// The internet assure me this should break out of the loop only, but I am sceptical
				}
			}

			err := d.saveData(entries)
			if err != nil {
				req.ActionSuccess <- Result{Success: false, Error: err}
			}
			req.ActionSuccess <- Result{Success: true, Error: nil}
		case "delete":
			taskString := "all"
			if task != "" {
				taskString = task
			}
			statusString := "all"
			if status != "" {
				statusString = status
			}
			fmt.Printf("Deleting task %s, status %s\n", taskString, statusString)

			// Use a single pass to filter entries in-place.
			for i := 0; i < len(entries); {
				entry := entries[i]
				// Conditions for deletion
				if (task == "all" && status == "") ||
					(task == "all" && status != "" && entry.Status == status) ||
					(task != "all" && entry.Task == task) {
					// Remove the entry by swapping with the last element and slicing off the last item
					// Gotta confess, this was pure ChatGPT but it's pretty nifty
					entries[i] = entries[len(entries)-1]
					entries = entries[:len(entries)-1]
				} else {
					i++ // Move to the next item only if no deletion happened - if there was deletion there's a new entry to check in slot i
				}
			}

			err := d.saveData(entries)
			if err != nil {
				req.ActionSuccess <- Result{Success: false, Error: err}
			}
			req.ActionSuccess <- Result{Success: true, Error: nil}
		}
	}
	return nil
}

func (d *ConcurrentDatasource) getData() ([]ListEntry, error) {
	var todoList TodoList
	todoListData, readErr := JsonBytesFromFile(d.Filepath)
	if readErr != nil {
		return todoList.Entries, readErr
	}

	jsonErr := json.Unmarshal(todoListData, &todoList)
	if jsonErr != nil {
		return todoList.Entries, readErr
	}

	return todoList.Entries, nil
}

func (d *ConcurrentDatasource) saveData(entries []ListEntry) error {
	byteCount, err := ListToFile(d.Filepath, entries...)
	if err != nil {
		return err
	}
	fmt.Printf("Saved %s, %d bytes\n", d.Filepath, byteCount)
	return nil
}

func (d *ConcurrentDatasource) Create(task string, status string) error {
	success := make(chan Result)
	d.Requests <- DataRequest{
		CrudRequest: CrudRequest{
			Action: "create",
			Task:   task,
			Status: status,
		},
		ActionSuccess: success,
	}
	result := <-success
	return result.Error
}

func (d *ConcurrentDatasource) Read(task string, status string) ([]ListEntry, error) {
	success := make(chan Result)
	listToView := make(chan []ListEntry)
	d.Requests <- DataRequest{
		CrudRequest: CrudRequest{
			Action: "read",
			Task:   task,
			Status: status,
		},
		ActionSuccess: success,
		ListToView:    listToView,
	}
	result := <-success
	list := <-listToView

	return list, result.Error
}

func (d *ConcurrentDatasource) Update(task string, status string) error {
	success := make(chan Result)
	d.Requests <- DataRequest{
		CrudRequest: CrudRequest{
			Action: "update",
			Task:   task,
			Status: status,
		},
		ActionSuccess: success,
	}
	result := <-success

	return result.Error
}

func (d *ConcurrentDatasource) Delete(task string, status string) error {
	success := make(chan Result)
	d.Requests <- DataRequest{
		CrudRequest: CrudRequest{
			Action: "delete",
			Task:   task,
			Status: status,
		},
		ActionSuccess: success,
	}
	result := <-success

	return result.Error
}
