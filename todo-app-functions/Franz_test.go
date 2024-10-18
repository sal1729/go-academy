package todo_app_functions

import (
	"testing"
)

func TestCrudRequest_SetActionAndFilter(t *testing.T) {
	tests := []struct {
		name          string
		initialAction string
		action        string
		task          string
		status        string
		expectError   bool
		errorMessage  string
	}{
		{"Set successfully - create task status", "", "create", "Task One", "To Do", false, ""},
		{"Err successfully - create task", "", "create", "Task One", "", true, "cannot create or update without status"},
		{"Err successfully - create all status", "", "create", "all", "To Do", true, "cannot create 'all' - reserved keyword"},
		{"Err successfully - create all", "", "create", "all", "To Do", true, "cannot create 'all' - reserved keyword"},
		{"Err successfully - create", "", "create", "", "", true, "task cannot be empty"},

		{"Err successfully - list task status", "", "list", "Task One", "To Do", true, "cannot filter individual task by status when using list or delete"},
		{"Set successfully - list task", "", "list", "Task One", "", false, ""},
		{"Set successfully - list all status", "", "list", "all", "To Do", false, ""},
		{"Set successfully - list all", "", "list", "all", "", false, ""},
		{"Err successfully - list", "", "list", "", "", true, "task cannot be empty"},

		{"Set successfully - update task status", "", "update", "Task One", "In Progress", false, ""},
		{"Err successfully - update task", "", "update", "Task One", "", true, "cannot create or update without status"},
		{"Err successfully - update all status", "", "update", "all", "To Do", true, "update status must be of the form 'status*TO*newStatus"},
		{"Set successfully - update all status*TO*newStatus", "", "update", "all", "To Do*TO*In Progress", false, ""},
		{"Err successfully - update all", "", "update", "all", "", true, "cannot create or update without status"},
		{"Err successfully - update", "", "update", "", "", true, "task cannot be empty"},

		{"Err successfully - delete task status", "", "delete", "Task One", "To Do", true, "cannot filter individual task by status when using list or delete"},
		{"Set successfully - delete task", "", "delete", "Task One", "", false, ""},
		{"Set successfully - delete all status", "", "delete", "all", "To Do", false, ""},
		{"Set successfully - delete all", "", "delete", "all", "", false, ""},
		{"Err successfully - delete", "", "delete", "", "", true, "task cannot be empty"},

		{"Err successfully - action already set", "create", "update", "Task One", "To Do", true, "action (create) is already set"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := CrudRequest{Action: tt.initialAction}
			err := output.SetActionAndFilter(tt.action, tt.task, tt.status)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error '%s' but got none", tt.errorMessage)
				} else if err.Error() != tt.errorMessage {
					t.Errorf("expected error '%s', got '%s'", tt.errorMessage, err.Error())
				}
			} else {
				expected := CrudRequest{
					Action: tt.action,
					Task:   tt.task,
					Status: tt.status,
				}
				if err != nil {
					t.Errorf("expected no error but got '%s'", err)
				}
				if output != expected {
					t.Errorf("expected output %s, but got %s", expected, output)
				}
			}
		})
	}
}

func TestGetCrudRequestFromUserInput(t *testing.T) {
	tests := []struct {
		name           string
		input          UserInput
		expectedAction string
		expectedTask   string
		expectedStatus string
		expectErr      bool
		errorMessage   string
	}{
		{
			name:           "Create task with status",
			input:          UserInput{Create: "Task1", Status: "To Do"},
			expectedAction: "create",
			expectedTask:   "Task1",
			expectedStatus: "To Do",
			expectErr:      false,
		},
		{
			name:           "List task with no status",
			input:          UserInput{List: "Task2"},
			expectedAction: "list",
			expectedTask:   "Task2",
			expectedStatus: "",
			expectErr:      false,
		},
		{
			name:           "Update task with status",
			input:          UserInput{Update: "Task3", Status: "In Progress"},
			expectedAction: "update",
			expectedTask:   "Task3",
			expectedStatus: "In Progress",
			expectErr:      false,
		},
		{
			name:           "Delete task with no status",
			input:          UserInput{Delete: "Task4"},
			expectedAction: "delete",
			expectedTask:   "Task4",
			expectedStatus: "",
			expectErr:      false,
		},
		{
			name:         "Too many actions",
			input:        UserInput{Create: "Task6", Update: "Task5", Status: "Done"},
			expectErr:    true,
			errorMessage: "action (create) is already set",
		},
		{
			name:         "No action",
			input:        UserInput{Status: "Done"},
			expectErr:    true,
			errorMessage: "no action specified",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetCrudRequestFromUserInput(tt.input)

			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
				// Check if the error message matches the expected error
				if err.Error() != tt.errorMessage {
					t.Errorf("Expected error message '%s', but got '%s'", tt.errorMessage, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if actual.Action != tt.expectedAction {
				t.Errorf("Expected action %s, got %s", tt.expectedAction, actual.Action)
			}
			if actual.Task != tt.expectedTask {
				t.Errorf("Expected task %s, got %s", tt.expectedTask, actual.Task)
			}
			if actual.Status != tt.expectedStatus {
				t.Errorf("Expected status %s, got %s", tt.expectedStatus, actual.Status)
			}
		})
	}
}
