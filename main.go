package main

import (
	"fmt"
	"sourced_go/repository"
	"sourced_go/todos"
)

func main() {
	// Initialize the repository
	repo := repository.NewRepository()

	// Create a new ToDo instance
	todo := todos.NewToDo()

	// Register event listeners
	todo.On("ToDoInitialized", func(data interface{}) {
		eventData := data.(map[string]string) // Cast directly to map[string]string
		fmt.Printf("Task initialized: %v\n", eventData)
	})

	todo.On("ToDoCompleted", func(data interface{}) {
		eventData := data.(map[string]string) // Cast directly to map[string]string
		fmt.Printf("Task completed: %v\n", eventData)
	})

	// Initialize and complete the task
	todo.Initialize("todo-id-1", "0x000000test0000000", "The only task that matters")
	todo.Complete()

	// Commit the ToDo to the repository, which will store it and emit events
	repo.Commit(todo.Entity)

	// Use FindByID to look up the task we just committed
	found := repo.FindByID("todo-id-1")
	if found != nil {
		fmt.Printf("Found ToDo: %v with history: %v\n", found.ID, found.Commands)
	} else {
		fmt.Println("ToDo not found.")
	}

	// Try to find a non-existing ID
	found = repo.FindByID("non-existing-id")
	if found == nil {
		fmt.Println("ToDo not found for non-existing ID.")
	}
}
