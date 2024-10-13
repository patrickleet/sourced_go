package repository

import (
	"sourced_go/entity"
	"sync"
)

// Repository manages entities and commits events
type Repository struct {
	storage map[string]*entity.Entity
	mu      sync.Mutex
}

// NewRepository initializes a new repository with an in-memory storage
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string]*entity.Entity),
	}
}

// Commit stores the entity and commits its events
func (r *Repository) Commit(e *entity.Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Store the entity in the in-memory map using its ID as the key
	r.storage[e.ID] = e

	// Call the entity's EmitQueuedEvents method to emit and clear enqueued events
	e.EmitQueuedEvents()
}

// FindByID retrieves an entity by its ID from the in-memory storage
func (r *Repository) FindByID(id string) *entity.Entity {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Lookup the entity by its ID in the map
	if entity, exists := r.storage[id]; exists {
		return entity
	}
	return nil
}
