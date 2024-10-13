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

	// This will NOT be triggered in this example, because Complete() is not called
	// on this object `todo` - it is called later on by the rehydrated object `rehydratedTodo`
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
		fmt.Println("Found ToDo:", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)
	} else {
		fmt.Println("ToDo not found.")
	}

	// This one will trigger - I point this out to say be careful with
	// binding events to objects and then rehydrating them...
	// I don't see any reason you'd ever need to create an object and then
	// rehydrate it in the same handler, but it's something to be aware of.
	rehydratedTodo.On("ToDoCompleted", func(data interface{}) {
		eventData := data.(map[string]string)
		fmt.Printf("Task completed: %v\n", eventData)
	})

	rehydratedTodo.Complete()
	fmt.Println("Reyhdrated ToDo Completed:", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

	todoRepo.Commit(rehydratedTodo)

	// Retrieve the ToDo by ID again
	todo3 := todoRepo.FindByID("todo-id-1")
	if todo3 != nil {
		fmt.Println("Found ToDo again:", todo3.ID, todo3.Task, todo3.Completed)
	} else {
		fmt.Println("ToDo not found.")
	}
}
