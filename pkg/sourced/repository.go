package sourced

import (
	"sync"
)

// Repository manages command logs for event-sourced entities
type Repository struct {
	storage map[string][]CommandRecord
	mu      sync.Mutex
}

// NewRepository initializes a new repository
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string][]CommandRecord),
	}
}

// FindByID retrieves an entity by ID, returning the raw generic Entity
func (r *Repository) FindByID(id string) *Entity {
	r.mu.Lock()
	defer r.mu.Unlock()

	if commands, exists := r.storage[id]; exists {
		e := &Entity{ID: id}
		e.Commands = commands // Load commands, but don't rehydrate yet

		// Ensure the EventEmitter is initialized during rehydration
		if e.EventEmitter == nil {
			e.EventEmitter = NewEventEmitter() // Initialize it properly
		}
		return e
	}
	return nil
}

// Commit stores the commands executed on the entity
func (r *Repository) Commit(e *Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the command log (list of commands)
	r.storage[e.ID] = e.Commands

	// Emit all queued events
	e.EmitQueuedEvents()
}
