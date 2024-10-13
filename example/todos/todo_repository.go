package todos

import (
	"sourced_go/pkg/sourced"
)

// ToDoRepository provides domain-specific behavior for ToDo entities
type ToDoRepository struct {
	*sourced.Repository // Extend the generic repository
}

// NewToDoRepository initializes a new ToDoRepository
func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		Repository: sourced.NewRepository(),
	}
}

// FindByID retrieves and rehydrates a ToDo by ID
func (r *ToDoRepository) FindByID(id string) *ToDo {
	// Get the generic entity from the base repository
	rehydratedEntity := r.Repository.FindByID(id)
	if rehydratedEntity == nil {
		return nil
	}

	// Now rehydrate the specific ToDo from the entity's events
	todo := &ToDo{Entity: rehydratedEntity}
	todo.Replaying = true // Set replaying flag to prevent digesting during rehydration

	for _, cmd := range rehydratedEntity.Events {
		todo.ReplayEvent(cmd)
	}

	todo.Replaying = false // Reset replaying flag
	return todo
}

// Commit commits the ToDo entity to the repository
func (r *ToDoRepository) Commit(t *ToDo) {
	r.Repository.Commit(t.Entity) // Use the generic repository to commit
}
