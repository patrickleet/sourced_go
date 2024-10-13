package main

import (
	"fmt"
	"sourced_go/todos"
)

func main() {
	repository := todos.NewToDoRepository()
	todo := todos.NewToDo()

	// Register event listeners
	repository.On("ToDoInitialized", func(data interface{}) {
		event := data.(todos.ToDoInitialized)
		fmt.Printf("Task initialized: %v\n", event.Task)
	})

	repository.On("ToDoCompleted", func(data interface{}) {
		event := data.(todos.ToDoCompleted)
		fmt.Printf("Task completed: %v\n", event.ID)
	})

	// Initialize and complete a task
	todo.Initialize("todo-id-1", "0x000000test0000000", "The only task that matters")
	todo.Complete()

	// Save the task to the repository (this triggers the commit and emits the events)
	repository.SaveToDo(todo)
}
