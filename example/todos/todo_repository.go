package todos

import (
	"sourced_go/pkg/sourced"
	"sync"
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
	var todos []*ToDo
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			entity := r.Repository.Get(id)
			if entity != nil {
				todo := &ToDo{Entity: entity}
				todo.Replaying = true
				for _, event := range entity.Events {
					todo.ReplayEvent(event)
				}
				todo.Replaying = false

				mu.Lock()
				todos = append(todos, todo)
				mu.Unlock()
			}
		}(id)
	}

	wg.Wait()
	return todos
}

// Commit commits the ToDo entity to the repository
func (r *ToDoRepository) Commit(t *ToDo) {
	r.Repository.Commit(t.Entity) // Use the generic repository to commit
}

func (r *ToDoRepository) CommitAll(todos []*ToDo) {
	entities := make([]*sourced.Entity, len(todos))
	var wg sync.WaitGroup

	for i := range todos {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			entities[i] = todos[i].Entity
		}(i)
	}

	wg.Wait()
	r.Repository.CommitAll(entities)

	// Emit queued events in parallel
	for _, todo := range todos {
		wg.Add(1)
		go func(todo *ToDo) {
			defer wg.Done()
			todo.EmitQueuedEvents()
		}(todo)
	}

	wg.Wait()
}
