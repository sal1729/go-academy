package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	todo "todo_app_functions"
)

func worker(id int, tasks <-chan todo.ListEntry, wg *sync.WaitGroup, listLock *sync.Mutex) {
	defer wg.Done()
	for t := range tasks {
		// Random sleep to simulate processing time
		// When there was no sleeping worker 2 did all the work
		sleepDuration := rand.Intn(500) // Sleep for 0 to 500 milliseconds
		time.Sleep(time.Duration(sleepDuration) * time.Millisecond)

		listLock.Lock()
		fmt.Printf("Worker %d slept for %d ms: %s\n", id, sleepDuration, t)
		//fmt.Printf("Worker %d: %s\n", id, t)
		listLock.Unlock()
	}
}

func main() {
	todoList := []todo.ListEntry{
		{Task: "Feed the floor", Status: "To Do"},
		{Task: "Sweep the dishes", Status: "Done"},
		{Task: "Rock the rug", Status: "Blocked"},
		{Task: "Scrub the fishes", Status: "In Progress"},
		{Task: "Vacuum the lawn", Status: "Done"},
		{Task: "Bathe the mat", Status: "To Do"},
		{Task: "Mop the baby", Status: "In Progress"},
		{Task: "Mow the cat", Status: "In Progress"},
		{Task: "Stop! Look!", Status: "Done"},
		{Task: "Buy the book", Status: "Blocked"},
	}

	tasks := make(chan todo.ListEntry, len(todoList))
	listLock := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	// Create worker pool
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, tasks, wg, listLock)
	}

	// Send tasks to channel, then close
	for t := range todoList {
		tasks <- todoList[t]
	}
	close(tasks)

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("Done")
}
