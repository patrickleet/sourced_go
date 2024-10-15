package sourced

import "sync"

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
	var mu sync.Mutex // To protect shared data
	var wg sync.WaitGroup

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			if entity := r.Get(id); entity != nil {
				mu.Lock()
				entities = append(entities, entity)
				mu.Unlock()
			}
		}(id)
	}

	wg.Wait()
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
	var wg sync.WaitGroup

	for _, entity := range entities {
		wg.Add(1)
		go func(entity *Entity) {
			defer wg.Done()
			r.storage[entity.ID] = entity.Events
		}(entity)
	}

	// Emit all queued events for each entity
	for _, entity := range entities {
		entity.EmitQueuedEvents()
	}

	wg.Wait()
}
