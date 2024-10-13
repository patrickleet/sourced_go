package todos

import (
	"sourced_go/pkg/sourced"
)

// ToDo represents the ToDo model
type ToDo struct {
	*sourced.Entity
	UserId    string
	Task      string
	Completed bool
}

// NewToDo creates a new ToDo instance
func NewToDo() *ToDo {
	return &ToDo{
		Entity: sourced.NewEntity(),
	}
}

// Initialize sets up the ToDo task and enqueues the initialized event
func (t *ToDo) Initialize(id, userId, task string) {
	t.ID = id
	t.UserId = userId
	t.Task = task
	t.Completed = false

	t.Digest("Initialize", id, userId, task)
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
func (t *ToDo) ReplayCommand(cmd sourced.CommandRecord) {
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
		"UserId":    t.UserId,
		"Task":      t.Task,
		"Completed": t.Completed,
	}
}
