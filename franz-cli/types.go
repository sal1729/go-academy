package main

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

type TodoItem struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}
