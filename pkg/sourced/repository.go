package sourced

import (
	"sync"
)

// Repository manages events logs for event-sourced entities
type Repository struct {
	storage map[string][]EventRecord
	mu      sync.Mutex
}

// NewRepository initializes a new repository
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string][]EventRecord),
	}
}

// FindByID retrieves an entity by ID, returning the raw generic Entity
func (r *Repository) FindByID(id string) *Entity {
	r.mu.Lock()
	defer r.mu.Unlock()

	if events, exists := r.storage[id]; exists {
		e := &Entity{ID: id}
		e.Events = events // Load events, but don't rehydrate yet

		// Ensure the EventEmitter is initialized during rehydration
		if e.EventEmitter == nil {
			e.EventEmitter = NewEventEmitter() // Initialize it properly
		}
		return e
	}
	return nil
}

// Commit stores the events executed on the entity
func (r *Repository) Commit(e *Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the events log (list of events)
	r.storage[e.ID] = e.Events

	// Emit all queued events
	e.EmitQueuedEvents()
}
