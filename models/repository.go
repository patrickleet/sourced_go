package models

import "sync"

// Repository manages entities and emits events after successful commits
type Repository struct {
	storage      map[string]*Entity // In-memory storage for entities
	EventEmitter *EventEmitter      // EventEmitter for emitting events after commit
	mu           sync.Mutex         // Mutex for thread safety
}

// NewRepository creates a new repository instance
func NewRepository() *Repository {
	return &Repository{
		storage:      make(map[string]*Entity),
		EventEmitter: NewEventEmitter(), // Initialize the event emitter
	}
}

// Save stores the entity and commits events
func (r *Repository) Save(entity *Entity) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Simulate storing the entity (in-memory in this case)
	r.storage[entity.ID] = entity

	// Commit the entity, which will trigger event emission
	r.Commit(entity)
}

// Commit emits the events after a successful commit
func (r *Repository) Commit(entity *Entity) {
	// Emit events only after a successful "commit"
	for _, event := range entity.EventsToEmit {
		r.EventEmitter.Emit(event.EventType(), event)
	}

	// Clear the events after emitting
	entity.Commit()
}

// FindByID retrieves an entity from the repository by its ID
func (r *Repository) FindByID(id string) *Entity {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.storage[id]
}
