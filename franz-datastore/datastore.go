package franz_datastore

import (
	"errors"
	"fmt"
)

type AccessRequest struct {
	CrudRequest   CrudRequest
	Result        chan<- []ListItem
	ActionSuccess chan<- err
}

// This doesn't feel right, I'm going to end up circumnavigating the api here 🤕
func datastore(requests <-chan AccessRequest) {
	todoList, err := GetData()
	if err != nil {
		panic(err)
	}

	for req := range requests {
		switch req.CrudRequest.Action {
		case "create":
		//
		case "read":
		//
		case "update":
		//
		case "delete":
		//
		default:
			errorString := fmt.Sprintln("invalid action:", req.CrudRequest.Action)
			fmt.Println(errorString)
			req.ActionSuccess <- errors.New(errorString)
		}
	}
}
