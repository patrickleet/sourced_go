package main

import (
	"fmt"
	"sourced_go/example/todos"
)

func main() {
	// Initialize the repository
	repo := todos.NewToDoRepository()

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
	repo.Commit(todo)

	// Retrieve the ToDo by ID
	rehydratedTodo := repo.Get("todo-id-1")

	if rehydratedTodo != nil {
		fmt.Println("Found ToDo", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

		// Re-bind event listeners to the rehydrated ToDo
		rehydratedTodo.On("ToDoCompleted", func(data interface{}) {
			todoInstance := data.(*todos.ToDo) // Cast the entity reference
			snapshot := todoInstance.Snapshot()
			fmt.Println("on ToDoCompleted", snapshot)
		})

		// Complete the rehydrated task
		rehydratedTodo.Complete()
		fmt.Println("Reyhdrated ToDo Completed", rehydratedTodo.ID, rehydratedTodo.Task, rehydratedTodo.Completed)

		repo.Commit(rehydratedTodo)
	} else {
		fmt.Println("ToDo not found.")
	}

	// Retrieve the ToDo by ID again
	todo3 := repo.Get("todo-id-1")
	if todo3 != nil {
		fmt.Println("Found ToDo again:", todo3.ID, todo3.Task, todo3.Completed)
	} else {
		fmt.Println("ToDo not found.")
	}

	// Initialize multiple ToDos
	all1 := todos.NewToDo()
	all1.Initialize("1", "user1", "Buy Sauna")

	all2 := todos.NewToDo()
	all2.Initialize("2", "user2", "Chew bubblegum")

	// Commit multiple ToDos to the repository
	repo.CommitAll([]*todos.ToDo{all1, all2})

	// Get all ToDos from the repository
	allTodos := repo.GetAll([]string{"1", "2"})

	if len(allTodos) > 0 {
		fmt.Println("All Todos:", allTodos)
		for _, todo := range allTodos {
			fmt.Println("ToDo:", todo.ID, todo.Task, todo.Completed)
		}
	} else {
		fmt.Println("No Todos found")
	}
}
