package todos

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
	r.Commit(todo.Entity)
}
