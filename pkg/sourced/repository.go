package sourced

// Repository manages events logs for event-sourced entities
type Repository struct {
	storage map[string][]EventRecord
}

// NewRepository initializes a new repository
func NewRepository() *Repository {
	return &Repository{
		storage: make(map[string][]EventRecord),
	}
}

// Get retrieves an entity by ID, returning the raw generic Entity
func (r *Repository) Get(id string) *Entity {
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

func (r *Repository) GetAll(ids []string) []*Entity {
	var entities []*Entity
	for _, id := range ids {
		if entity := r.Get(id); entity != nil {
			entities = append(entities, entity)
		}
	}
	return entities
}

// Commit stores the events executed on the entity
func (r *Repository) Commit(e *Entity) {
	// Store the events log (list of events)
	r.storage[e.ID] = e.Events

	// Emit all queued events
	e.EmitQueuedEvents()
}

func (r *Repository) CommitAll(entities []*Entity) {
	for _, entity := range entities {
		// Store the event log for each entity
		r.storage[entity.ID] = entity.Events
	}

	// After storing, emit events for each entity
	for _, entity := range entities {
		entity.EmitQueuedEvents()
	}
}
