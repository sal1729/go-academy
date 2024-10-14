package todo_app_functions

type ListEntry struct {
	Task string `json:"task"`
}

func (e ListEntry) String() string {
	return e.Task
}

type TodoList struct {
	Entries []ListEntry `json:"entries"`
}
