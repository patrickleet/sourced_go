package models

import (
	"sourced_go/repository"
)

// ToDoRepository inherits from Repository and exposes event-emitting behavior
type ToDoRepository struct {
	*repository.Repository // Inherits Repository and its behavior, including event handling
}

// NewToDoRepository initializes a new ToDoRepository
func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		Repository: repository.NewRepository(), // Initialize the base repository
	}
}

// SaveToDo saves the ToDo entity and commits events
func (r *ToDoRepository) SaveToDo(todo *ToDo) {
	r.Save(todo.Entity)
}

// On registers an event listener in the underlying EventEmitter
func (r *ToDoRepository) On(event string, listener func(data interface{})) {
	r.EventEmitter.On(event, listener)
}
