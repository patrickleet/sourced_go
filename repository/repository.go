package repository

import (
	"sourced_go/entity"
	"sync"
)

// Repository manages command logs for event-sourced entities
type Repository struct {
	storage map[string][]entity.CommandRecord
	mu      sync.Mutex
}

// NewRepository initializes a new repository
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string][]entity.CommandRecord),
	}
}

// FindByID retrieves an entity by ID, returning the raw generic Entity
func (r *Repository) FindByID(id string) *entity.Entity {
	r.mu.Lock()
	defer r.mu.Unlock()

	if commands, exists := r.storage[id]; exists {
		e := &entity.Entity{ID: id}
		e.Commands = commands // Load commands, but don't rehydrate yet
		return e
	}
	return nil
}

// Commit stores the commands executed on the entity
func (r *Repository) Commit(e *entity.Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the command log (list of commands)
	r.storage[e.ID] = e.Commands

	// Emit all queued events
	e.EmitQueuedEvents()
}
