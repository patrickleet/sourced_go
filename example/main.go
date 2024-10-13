package main

import (
	"fmt"
	"sourced_go/example/todos"
)

func main() {
	// Initialize the repository
	todoRepo := todos.NewToDoRepository()

	// Create a new ToDo instance
	todo := todos.NewToDo()

	// Register event listeners
	todo.On("ToDoInitialized", func(data interface{}) {
		todoInstance := data.(*todos.ToDo)
		fmt.Printf("Task initialized: ID=%v, Task=%v, Completed=%v\n", todoInstance.ID, todoInstance.Task, todoInstance.Completed)
	})

	// This will NOT be triggered in this example, because Complete() is not called
	// on this object `todo` - it is called later on by the rehydrated object `rehydratedTodo`
	todo.On("ToDoCompleted", func(data interface{}) {
		todoInstance := data.(*todos.ToDo) // Cast the entity reference
		fmt.Printf("Task completed: ID=%v, Task=%v, Completed=%v\n", todoInstance.ID, todoInstance.Task, todoInstance.Completed)
	})

	// Initialize and complete the task
	todo.Initialize("todo-id-1", "patrickleet", "Make sourced_go")

	// Commit the ToDo to the repository
	todoRepo.Commit(todo)

	// Retrieve the ToDo by ID
	rehydratedTodo := todoRepo.FindByID("todo-id-1")

	if rehydratedTodo != nil {
		fmt.Println("Found ToDo", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

		// Re-bind event listeners to the rehydrated ToDo
		rehydratedTodo.On("ToDoCompleted", func(data interface{}) {
			todoInstance := data.(*todos.ToDo) // Cast the entity reference
			snapshot := todoInstance.Snapshot()
			fmt.Println("Rehydrated Task completed:", snapshot)
		})

		// Complete the rehydrated task
		rehydratedTodo.Complete()
		fmt.Println("Reyhdrated ToDo Completed", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

		todoRepo.Commit(rehydratedTodo)
	} else {
		fmt.Println("ToDo not found.")
	}

	// Retrieve the ToDo by ID again
	todo3 := todoRepo.FindByID("todo-id-1")
	if todo3 != nil {
		fmt.Println("Found ToDo again:", todo3.ID, todo3.Task, todo3.Completed)
	} else {
		fmt.Println("ToDo not found.")
	}
}
