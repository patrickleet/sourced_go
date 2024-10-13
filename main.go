package main

import (
	"fmt"
	"sourced_go/todos"
)

func main() {
	// Initialize the repository
	todoRepo := todos.NewToDoRepository()

	// Create a new ToDo instance
	todo := todos.NewToDo()

	// Register event listeners
	todo.On("ToDoInitialized", func(data interface{}) {
		eventData := data.(map[string]string)
		fmt.Printf("Task initialized: %v\n", eventData)
	})

	todo.On("ToDoCompleted", func(data interface{}) {
		eventData := data.(map[string]string)
		fmt.Printf("Task completed: %v\n", eventData)
	})

	// Initialize and complete the task
	todo.Initialize("todo-id-1", "patrickleet", "Make sourced_go")

	// Commit the ToDo to the repository
	todoRepo.Commit(todo)

	// Retrieve the ToDo by ID
	rehydratedTodo := todoRepo.FindByID("todo-id-1")

	if rehydratedTodo != nil {
		fmt.Println("Found ToDo", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)
	} else {
		fmt.Println("ToDo not found.")
	}

	rehydratedTodo.Complete()
	fmt.Println("Reyhdrated ToDo Completed", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

	// todoRepo.Commit(rehydratedTodo)
}
