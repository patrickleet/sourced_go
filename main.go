package main

import (
	"fmt"
	"sourced_go/todos"
)

func main() {
	// Initialize repository
	repository := todos.NewToDoRepository()

	// Create a new ToDo instance and register event listeners
	todo := todos.NewToDo()
	todo.On("ToDoInitialized", func(data interface{}) {
		event := data.(todos.ToDoInitialized)
		fmt.Printf("Task initialized: %v\n", event.Task)
	})

	todo.On("ToDoCompleted", func(data interface{}) {
		event := data.(todos.ToDoCompleted)
		fmt.Printf("Task completed: %v\n", event.ID)
	})

	// Initialize and complete the task (enqueue events for later)
	todo.Initialize("todo-id-1", "0x000000test0000000", "The only task that matters")
	todo.Complete()

	// Now, commit the ToDo to the repository (this triggers commit and emission)
	repository.Commit(todo.Entity)
}
