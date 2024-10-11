package todo_app_functions

type ListEntry struct {
	Task string
}

func (e ListEntry) String() string {
	return e.Task
}
