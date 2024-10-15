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

// Get retrieves and rehydrates a ToDo by ID
func (r *ToDoRepository) Get(id string) *ToDo {
	// Get the generic entity from the base repository
	rehydratedEntity := r.Repository.Get(id)
	if rehydratedEntity == nil {
		return nil
	}

	// Now rehydrate the specific ToDo from the entity's events
	todo := &ToDo{Entity: rehydratedEntity}
	todo.Replaying = true // Set replaying flag to prevent digesting during rehydration

	for _, event := range rehydratedEntity.Events {
		todo.ReplayEvent(event)
	}

	todo.Replaying = false // Reset replaying flag
	return todo
}

func (r *ToDoRepository) GetAll(ids []string) []*ToDo {
	entities := r.Repository.GetAll(ids)
	todos := make([]*ToDo, len(entities))

	for i, entity := range entities {
		todo := &ToDo{Entity: entity}
		todo.Replaying = true
		for _, event := range entity.Events {
			todo.ReplayEvent(event)
		}
		todo.Replaying = false
		todos[i] = todo
	}

	return todos
}

// Commit commits the ToDo entity to the repository
func (r *ToDoRepository) Commit(t *ToDo) {
	r.Repository.Commit(t.Entity) // Use the generic repository to commit
}

func (r *ToDoRepository) CommitAll(todos []*ToDo) {
	entities := make([]*sourced.Entity, len(todos))
	for i, todo := range todos {
		entities[i] = todo.Entity
	}

	r.Repository.CommitAll(entities) // Commit all entities at once

	// Emit queued events for each ToDo after committing
	for _, todo := range todos {
		todo.EmitQueuedEvents()
	}
}
