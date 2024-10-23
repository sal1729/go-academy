package franz

type CrudRequest struct {
	Action string `json:"action"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type ListItem struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}
