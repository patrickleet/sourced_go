package main

import (
	"fmt"
	"sourced_go/models"
)

func main() {
	repository := models.NewToDoRepository()
	todo := models.NewToDo()

	// Register event listeners
	repository.On("ToDoInitialized", func(data interface{}) {
		event := data.(models.ToDoInitialized)
		fmt.Printf("Task initialized: %v\n", event.Task)
	})

	repository.On("ToDoCompleted", func(data interface{}) {
		event := data.(models.ToDoCompleted)
		fmt.Printf("Task completed: %v\n", event.ID)
	})

	// Initialize and complete a task
	todo.Initialize("todo-id-1", "0x000000test0000000", "The only task that matters")
	todo.Complete()

	// Save the task to the repository (this triggers the commit and emits the events)
	repository.SaveToDo(todo)
}
