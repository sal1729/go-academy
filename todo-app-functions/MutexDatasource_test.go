package todo_app_functions

import (
	"os"
	"reflect"
	"strconv"
	"sync"
	"testing"
)

func setupTestDatasource(t *testing.T, testname string) *LocalMuDatasource {
	// Use a temporary file for testing
	file, err := os.CreateTemp("", testname+".json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	return &LocalMuDatasource{
		Filepath: file.Name(),
	}
}

// Note, some missing test cases here - TestUpdateFilter and TestReadTask, but franz-datastore doesn't care about these.

func TestCreate(t *testing.T) {
	d := setupTestDatasource(t, "create-test")
	testTask := "Test Task"
	testStatus := "Test Status"

	// When we create a task
	expected := ListEntry{Task: testTask, Status: testStatus}
	received, err := d.Create(testTask, testStatus)
	if err != nil {
		t.Errorf("Failed to create item: %v", err)
	}

	// Then it has the expected form
	if !reflect.DeepEqual(received, expected) {
		t.Errorf("Incorrect created item. Expected: %s, Received: %s", expected, received)
	}

	// And the underlying data file has been updated
	items, err := d.Read("all", "") // TODO Doesn't feel right to assume this function is ok to use in a test?
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	if len(items) != 1 || items[0].Task != expected.Task {
		t.Errorf("Expected 1 entry with name %s, got %v", expected.Task, items)
	}
}

func TestUpdate(t *testing.T) {
	d := setupTestDatasource(t, "update-test")
	testTask := "Test Task"
	testStatus := "Test Status"
	newStatus := "New Status"

	// Given an item is created
	_, err := d.Create(testTask, testStatus)
	if err != nil {
		t.Errorf("Failed to create item: %v", err)
	}

	// When it is updated
	expected := ListEntry{Task: testTask, Status: newStatus}
	received, err := d.Update(testTask, newStatus)
	if err != nil {
		t.Errorf("Failed to update item: %v", err)
	}

	// Then it has the expected form
	if len(received) != 1 || !reflect.DeepEqual(received[0], expected) {
		t.Errorf("Incorrectly created item. Expected one task with status: %s, Received: %s", newStatus, received)
	}

	// And the underlying data file has been updated
	items, err := d.Read("all", "")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	if len(items) != 1 || items[0].Status != expected.Status {
		t.Errorf("Expected updated status to be %s, got %s", expected.Status, items[0].Status)
	}
}

func TestDeleteTask(t *testing.T) {
	d := setupTestDatasource(t, "delete-task-test")
	testTask := "Test Task"
	testStatus := "Test Status"

	// Given an item is created
	_, err := d.Create(testTask, testStatus)
	if err != nil {
		t.Errorf("Failed to create item: %v", err)
	}

	// When it is deleted
	deleted, err := d.Delete(testTask, "")
	if err != nil {
		t.Errorf("Failed to delete item: %v", err)
	}

	// Then the list of deleted items looks as expected
	if len(deleted) > 1 || deleted[0].Task != testTask {
		t.Errorf("Incorrect deleted item. Expected: %s, Received: %s", testTask, deleted)
	}

	// And the file has been updated
	items, err := d.Read("all", "")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %v", items)
	}
}

func TestDeleteAll(t *testing.T) {
	d := setupTestDatasource(t, "delete-all-test")
	testTask := "Test Task"
	testStatus := "Test Status"

	// Given multiple items are created
	for i := 1; i <= 5; i++ {
		_, err := d.Create(testTask+strconv.Itoa(i), testStatus)
		if err != nil {
			t.Errorf("Failed to create item: %v", err)
		}
	}

	// When delete is called with task "all"
	deleted, err := d.Delete("all", "")
	if err != nil {
		t.Errorf("Failed to delete items: %v", err)
	}

	// Then the list of deleted items contains all items
	if len(deleted) != 5 {
		t.Errorf("Incorrect count of deleted items. Expected: 5, Received: %d", len(deleted))
	}

	// And the file has been updated
	items, err := d.Read("all", "")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %v", items)
	}
}

func TestDeleteFilter(t *testing.T) {
	d := setupTestDatasource(t, "delete-filter-test")
	testTask := "Test Task"
	testStatus := "Test Status"

	// Given multiple items are created
	for i := 1; i <= 5; i++ {
		_, err := d.Create(testTask+strconv.Itoa(i), testStatus+" "+strconv.Itoa(i%2))
		if err != nil {
			t.Errorf("Failed to create item: %v", err)
		}
	}

	// When delete is called with non empty status
	deleted, err := d.Delete("all", "Test Status 0")
	if err != nil {
		t.Errorf("Failed to delete items: %v", err)
	}

	// Then the list of deleted items contains all expected items
	if len(deleted) != 2 && deleted[0].Status != "Test Status 0" {
		t.Errorf("Incorrect count of deleted items. Expected: 2, Received: %s", deleted)
	}

	// And the file has been updated
	items, err := d.Read("all", "")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	if len(items) != 3 || items[0].Status != "Test Status 1" {
		t.Errorf("Expected 3 items, got %v", items)
	}
}

func TestReadAll(t *testing.T) {
	d := setupTestDatasource(t, "read-all-test")

	testTask := "Test Task"
	testStatus := "Test Status"

	// Given multiple items are created
	for i := 1; i <= 5; i++ {
		_, err := d.Create(testTask+strconv.Itoa(i), testStatus+" "+strconv.Itoa(i%2))
		if err != nil {
			t.Errorf("Failed to create item: %v", err)
		}
	}

	// When they are read
	items, err := d.Read("all", "")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	// Then the list is correct
	if len(items) != 5 {
		t.Errorf("Expected 5 items, got %d", len(items))
	}
	for idx, i := range items {
		if reflect.DeepEqual(i, ListEntry{Task: testTask + strconv.Itoa(idx+1), Status: testStatus + " " + strconv.Itoa((idx+1)%2)}) {
			t.Errorf("Incorrect item at index %d: %v", idx+1, i)
		}
	}
}

func TestReadFilter(t *testing.T) {
	d := setupTestDatasource(t, "read-filter-test")

	testTask := "Test Task"
	testStatus := "Test Status"

	// Given multiple items are created
	for i := 1; i <= 5; i++ {
		_, err := d.Create(testTask+strconv.Itoa(i), testStatus+" "+strconv.Itoa(i%2))
		if err != nil {
			t.Errorf("Failed to create item: %v", err)
		}
	}

	// When some are read
	items, err := d.Read("all", "Test Status 0")
	if err != nil {
		t.Errorf("Failed to list items: %v", err)
	}

	// Then the list is correct
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
	for idx, i := range items {
		if !reflect.DeepEqual(i, ListEntry{Task: testTask + strconv.Itoa((idx*2)+2), Status: "Test Status 0"}) {
			t.Errorf("Incorrect item at index %d: %v", idx+1, i)
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	d := setupTestDatasource(t, "concurrent-access-test")
	testTask := "Test Task"
	testStatus := "Test Status"
	var wg sync.WaitGroup

	// Given many tasks
	for i := 0; i < 100; i++ {
		_, err := d.Create(testTask+strconv.Itoa(i), testStatus+" "+strconv.Itoa(i%5))
		if err != nil {
			t.Errorf("Failed to create item %d: %v", i, err)
		}
	}

	// When many writes (updates), reads and deletes take place
	for i := 0; i < 100; i++ {
		wg.Add(3)

		// Read
		go func() {
			defer wg.Done()
			_, err := d.Read("all", "")
			if err != nil {
				t.Errorf("Failed to list items: %v", err)
			}
		}()

		// Write
		go func(id int) {
			defer wg.Done()
			_, err := d.Update(testTask+strconv.Itoa(id), testStatus+" Updated")
			if err != nil {
				t.Errorf("Failed to update item %d: %v", id, err)
			}
		}(i)

		// Delete
		go func(id int) {
			defer wg.Done()
			_, err := d.Delete(testTask+strconv.Itoa(id), "")
			if err != nil {
				t.Errorf("Failed to delete item %d: %v", id, err)
			}
		}(i)
	}

	// And we wait for all goroutines to finish
	wg.Wait()

	// Then all is well and there is no data
	items, err := d.Read("all", "")
	if err != nil {
		t.Fatalf("Failed to list items: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("Expected 0 items after concurrent deletes, got %s", items)
	}
}
