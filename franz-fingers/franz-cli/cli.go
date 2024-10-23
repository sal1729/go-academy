// This is exercise 17 again
package main

import (
	"errors"
	"flag"
	"fmt"
	franz "franz_brain"
	"log"
)

var accessMethod = "raw" // Or change this to "goroutine"

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("franz: ")
	log.SetFlags(0)
	// Also set up a function for catching errors
	catchError := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	// Option flags
	var create, status, list, update, del string
	flag.StringVar(&create, "create", "", "task to create, must pass with -status")
	flag.StringVar(&status, "status", "", "status to create/read/update/delete, must be passed with other flags")
	flag.StringVar(&list, "list", "", "pass with all to list all tasks, with taskname to list single task, with all and -status to filter on status")
	flag.StringVar(&update, "update", "", "task to update, must pass with -status")
	flag.StringVar(&del, "delete", "", "pass with all to delete all tasks, with taskname to delete single task, with all and -status to delete by status")
	flag.Parse()

	// Convert user input to CrudRequest
	input := userInput{Create: create, Status: status, List: list, Update: update, Delete: del}
	request, err := getCrudRequestFromUserInput(input)
	catchError(err)

	if accessMethod == "raw" {
		// Get data
		todoList, err := franz.GetData()
		catchError(err)

		// Handle request
		err = handleRequest(&todoList, request)
		catchError(err)

		// Persist data
		if request.Action != "read" {
			err = franz.SaveData(&todoList)
			catchError(err)
		}

		log.Println("Request handled successfully:", request)
	}

	// Not sure what should happen here - should we be pinging off to another server?
	//if accessMethod == "api" {
	//	// Fire up request channel and datastore
	//	requests := make(chan franz.AccessRequest)
	//	go franz.Datastore(requests)
	//
	//	// Handle request via franz api
	//	franz.Api(w http.ResponseWriter, r *http.Request, requests chan franz.AccessRequest)
	//}

	if accessMethod == "goroutine" {
		// This just feels like a hacky version of the api
		// Fire up request channel and datastore
		requests := make(chan franz.AccessRequest)
		go franz.Datastore(requests)

		// Handle the request
		success := make(chan franz.SuccessResponse)
		response := make(chan []franz.ListItem)
		requests <- franz.AccessRequest{
			CrudRequest:   request,
			ActionSuccess: success,
			Result:        response,
		}

		complete := <-success
		if !complete.Success || complete.Error != nil {
			fmt.Printf("Request failed. Request: %s, Error: %v", request, complete.Error)
		} else if request.Action == "read" {
			list := <-response
			fmt.Println("List:", list)
		}

		// Shut down
		close(requests)
	}

}

func handleRequest(todoList *[]franz.ListItem, req franz.CrudRequest) error {
	// TODO Handle unhappy path where there was nothing to update/delete/list
	switch req.Action {
	case "create":
		newEntry, createErr := franz.CreateTask(todoList, req)
		if createErr != nil {
			return createErr
		}
		fmt.Println("Created entry:\n", newEntry)
	case "read":
		list, listErr := franz.ReadFromList(todoList, req)
		if listErr != nil {
			return listErr
		}
		fmt.Println("Listing entries:", list)
	case "update":
		updates, updateErr := franz.UpdateListItems(todoList, req)
		if updateErr != nil {
			return updateErr
		}
		fmt.Println("Updated entries:", updates)
	case "delete":
		deletions, deleteErr := franz.DeleteFromList(todoList, req)
		if deleteErr != nil {
			return deleteErr
		}
		fmt.Println("Deleted entries:", deletions)
	default:
		return fmt.Errorf("unknown action: %s", req.Action)
	}
	return nil
}

func setActionAndFilter(req *franz.CrudRequest, action string, task string, status string) error {
	if req.Action != "" {
		return fmt.Errorf("action (%s) is already set", req.Action)
	}
	req.Action = action
	req.Task = task
	req.Status = status

	// This might be overkill, as we use the same function to validate when we try to do the CRUD ops, but it allows us to fail faster
	err := franz.ValidateCrudRequest(*req)
	if err != nil {
		return err
	}

	return nil
}

type userInput struct {
	Create string
	Status string
	List   string
	Update string
	Delete string
}

func getCrudRequestFromUserInput(input userInput) (franz.CrudRequest, error) {
	var crudRequest franz.CrudRequest
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
		action = "read"
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
