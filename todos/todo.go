package todos

import "sourced_go/entity"

// ToDo represents the ToDo model
type ToDo struct {
	*entity.Entity
	Address   string
	Task      string
	Completed bool
	Removed   bool
}

// NewToDo creates a new ToDo instance
func NewToDo() *ToDo {
	return &ToDo{
		Entity: entity.NewEntity(),
	}
}

// Initialize sets up the ToDo task and enqueues the initialized event
func (t *ToDo) Initialize(id, address, task string) {
	t.ID = id
	t.Address = address
	t.Task = task
	t.Completed = false
	t.Removed = false
	t.DigestCommand("Initialize", id, address, task)

	// Enqueue the initialized event as a GenericEvent
	t.Enqueue(entity.GenericEvent{
		Type: "ToDoInitialized",
		Data: map[string]string{
			"ID":      id,
			"Address": address,
			"Task":    task,
		},
	})
}

// Complete marks the ToDo as completed and enqueues the completed event
func (t *ToDo) Complete() {
	if !t.Completed {
		t.Completed = true
		t.DigestCommand("Complete", t.ID)

		// Enqueue the completed event as a GenericEvent
		t.Enqueue(entity.GenericEvent{
			Type: "ToDoCompleted",
			Data: map[string]string{
				"ID": t.ID,
			},
		})
	}
}
