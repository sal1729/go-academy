package todo_app_functions

import "errors"

func GetCrudRequestFromUserInput(input UserInput) (CrudRequest, error) {
	var crudRequest CrudRequest
	var action string
	var task string
	var err error

	if input.Create != "" {
		action = "create"
		task = input.Create
		err = crudRequest.SetActionAndFilter(action, task, input.Status)
	}
	if input.List != "" {
		action = "list"
		task = input.List
		err = crudRequest.SetActionAndFilter(action, task, input.Status)
	}
	if input.Update != "" {
		action = "update"
		task = input.Update
		err = crudRequest.SetActionAndFilter(action, task, input.Status)
	}
	if input.Delete != "" {
		action = "delete"
		task = input.Delete
		err = crudRequest.SetActionAndFilter(action, task, input.Status)
	}

	if action == "" {
		err = errors.New("no action specified")
	}

	return crudRequest, err
}
