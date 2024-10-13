package todos

import "sourced_go/entity"

// ToDoInitialized event structure
type ToDoInitialized struct {
	ID      string
	Address string
	Task    string
}

func (e ToDoInitialized) EventType() string {
	return "ToDoInitialized"
}

// ToDoCompleted event structure
type ToDoCompleted struct {
	ID string
}

func (e ToDoCompleted) EventType() string {
	return "ToDoCompleted"
}

// ToDo represents the ToDo model, extending the core Entity
type ToDo struct {
	*entity.Entity
	Address   string
	Task      string
	Completed bool
	Removed   bool
}

// NewToDo creates a new ToDo entity
func NewToDo() *ToDo {
	return &ToDo{
		Entity: entity.NewEntity(), // No need to pass EventEmitter explicitly
	}
}

// Initialize sets up the ToDo and enqueues the "initialized" event
func (t *ToDo) Initialize(id, address, task string) {
	t.ID = id
	t.Address = address
	t.Task = task
	t.Completed = false
	t.Removed = false
	t.DigestCommand("Initialize", id, address, task)

	// Always enqueue the event for commit
	t.Enqueue(ToDoInitialized{id, address, task})
}

// Complete marks the ToDo as completed and enqueues the "completed" event
func (t *ToDo) Complete() {
	if !t.Completed {
		t.Completed = true
		t.DigestCommand("Complete", t.ID)

		// Always enqueue the event for commit
		t.Enqueue(ToDoCompleted{t.ID})
	}
}
