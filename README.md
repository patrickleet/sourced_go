# sourced_go

`sourced_go` is a Go-based implementation of an event-sourced architecture, inspired by the original [sourced](https://github.com/mateodelnorte/sourced) project by Matt Walters. Patrick Lee Scott (me), a contributor and maintainer of the original JavaScript/Typescript version, has brought these concepts to Go, extending and refactoring it for Go's ecosystem.

This repository demonstrates how to manage a `ToDo` task using event-sourcing, repositories, and event-driven architecture with Go. The core concept revolves around capturing state changes as commands (or events) and rehydrating entities by replaying those events.

## Prerequisites

- [Go](https://golang.org/dl/) installed (Go 1.16+ recommended)
- Basic understanding of event-sourcing principles

## Features

- **Event Sourcing**: All state changes to the `ToDo` are stored as events or commands.
- **In-Memory Repository**: Provides an easy way to store and retrieve `ToDo` tasks.
- **Event Emitters**: Register listeners and emit events when actions (e.g., task initialization or completion) occur.
- **Rehydration**: Rebuild the state of a `ToDo` by replaying its commands.

## Running the Example

See [`example/main.go`](https://github.com/patrickleet/sourced_go/blob/main/example/main.go). This shows how to use the repository and the event emitter with `ToDo` tasks.

To run:

```sh
go run example/main.go
```

### Key Concepts Explained

1. **ToDo Creation and Initialization**: 
   - A new `ToDo` is created, and event listeners are registered to respond to the `ToDoInitialized` and `ToDoCompleted` events.
   - The `Initialize` method sets up the task and triggers the `ToDoInitialized` event.

2. **Committing to Repository**: 
   - After initialization, the `ToDo` is committed to the repository, storing its events.

3. **Rehydrating the ToDo**: 
   - When a `ToDo` is retrieved from the repository (`FindByID`), the stored events are replayed, rebuilding the `ToDo`'s state.

4. **Re-binding Event Listeners**: 
   - After retrieving the rehydrated `ToDo`, event listeners must be re-bound for any subsequent actions (like `Complete()`).

5. **Re-completion**: 
   - After the rehydration, the `ToDo` is completed again, and the event is triggered.

## Key concepts

Sourced relies on the repository pattern to allow models to be clean and simple.

```go
// Get  repository
todoRepo := todos.NewToDoRepository()

// Create a new ToDo instance
todo := todos.NewToDo()

// Register event listeners
todo.On("ToDoInitialized", func(data interface{}) {
  todoInstance := data.(*todos.ToDo)
  fmt.Printf("Task initialized: ID=%v, Task=%v, Completed=%v\n", todoInstance.ID, todoInstance.Task, todoInstance.Completed)
})

// Initialize and complete the task
todo.Initialize("todo-id-1", "patrickleet", "Make sourced_go")

// Commit the ToDo to the repository
todoRepo.Commit(todo)
```

## Contributing

Feel free to open issues or submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License.
