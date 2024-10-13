package todos

import (
	"sourced_go/repository"
)

// ToDoRepository provides domain-specific behavior for ToDo entities
type ToDoRepository struct {
	*repository.Repository // Extend the generic repository
}

// NewToDoRepository initializes a new ToDoRepository
func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		Repository: repository.NewRepository(),
	}
}

// FindByID retrieves and rehydrates a ToDo by ID
func (r *ToDoRepository) FindByID(id string) *ToDo {
	// Get the generic entity from the base repository
	entity := r.Repository.FindByID(id)
	if entity == nil {
		return nil
	}

	// Now rehydrate the specific ToDo from the entity's commands
	todo := &ToDo{Entity: entity}
	for _, cmd := range entity.Commands {
		todo.ReplayCommand(cmd)
	}
	return todo
}

// Commit commits the ToDo entity to the repository
func (r *ToDoRepository) Commit(t *ToDo) {
	r.Repository.Commit(t.Entity) // Use the generic repository to commit
}
