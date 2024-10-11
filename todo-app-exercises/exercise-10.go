package main

import (
	"fmt"
	todo "todo_app_functions"
)

// Create a program using a variadic function to print a list of 10 things To Do.
func main() {
	todoList := []todo.ListEntry{
		{Task: "Feed the floor"},
		{Task: "Sweep the dishes"},
		{Task: "Rock the rug"},
		{Task: "Scrub the fishes"},
		{Task: "Vacuum the lawn"},
		{Task: "Bathe the mat"},
		{Task: "Mop the baby"},
		{Task: "Mow the cat"},
		{Task: "Stop! Look!"},
		{Task: "Buy the book"},
	}
	printToDoList(todoList...)
}

func printToDoList(list ...todo.ListEntry) {
	for _, entry := range list {
		fmt.Println(entry)
	}
}
