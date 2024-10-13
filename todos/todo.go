package todos

import (
	"sourced_go/entity"
)

// ToDo represents the ToDo model
type ToDo struct {
	*entity.Entity
	Address   string
	Task      string
	Completed bool
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

	t.Digest("Initialize", id, address, task)

	t.Enqueue("ToDoInitialized", t)
}

// Complete marks the ToDo as completed and enqueues the completed event
func (t *ToDo) Complete() {
	if !t.Completed {
		t.Completed = true
		t.Digest("Complete", t.ID)

		t.Enqueue("ToDoCompleted", t)
	}
}

// ReplayCommand replays the commands for the ToDo entity
func (t *ToDo) ReplayCommand(cmd entity.CommandRecord) {
	switch cmd.CommandName {
	case "Initialize":
		// Ensure that the correct arguments are passed and task is set
		t.Initialize(cmd.Args[0].(string), cmd.Args[1].(string), cmd.Args[2].(string))
	case "Complete":
		t.Complete()
	default:
		// Handle unknown commands if necessary
	}
}

func (t *ToDo) Snapshot() map[string]interface{} {
	return map[string]interface{}{
		"ID":        t.ID,
		"Address":   t.Address,
		"Task":      t.Task,
		"Completed": t.Completed,
	}
}
