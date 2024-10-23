package franz_datastore

import (
	"errors"
	"fmt"
)

type SuccessResponse struct {
	Success bool
	Error   error
}

type AccessRequest struct {
	CrudRequest   CrudRequest
	Result        chan<- []ListItem
	ActionSuccess chan<- SuccessResponse
}

func Datastore(requests <-chan AccessRequest) {
	todoList, err := GetData()
	if err != nil {
		panic(err)
	}

	for req := range requests {
		c := req.CrudRequest
		// This is an adaptation of the cli tool handleRequest function
		switch c.Action {
		case "create":
			newEntry, err := CreateTask(&todoList, c)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}
			err = SaveData(&todoList)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}

			fmt.Println("Created entry:\n", newEntry)
			req.ActionSuccess <- SuccessResponse{Success: true, Error: nil}
		case "read":
			list, err := ReadFromList(&todoList, c)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}

			fmt.Println("Listing entries:", list)
			req.ActionSuccess <- SuccessResponse{Success: true, Error: nil}
			req.Result <- list
		case "update":
			updates, err := UpdateListItems(&todoList, c)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}
			err = SaveData(&todoList)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}

			fmt.Println("Updated entries:", updates)
			req.ActionSuccess <- SuccessResponse{Success: true, Error: nil}
		case "delete":
			deletions, err := DeleteFromList(&todoList, c)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}
			err = SaveData(&todoList)
			if err != nil {
				req.ActionSuccess <- SuccessResponse{Success: false, Error: err}
				break
			}

			fmt.Println("Deleted entries:", deletions)
			req.ActionSuccess <- SuccessResponse{Success: true, Error: nil}
		default:
			errorString := fmt.Sprintln("invalid action:", c.Action)
			fmt.Println(errorString)
			req.ActionSuccess <- SuccessResponse{Success: false, Error: errors.New(errorString)}
		}
	}
}
