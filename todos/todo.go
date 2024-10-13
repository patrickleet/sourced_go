package todos

import (
	"fmt"
	"sourced_go/entity"
)

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

		t.Enqueue(entity.GenericEvent{
			Type: "ToDoCompleted",
			Data: map[string]string{
				"ID": t.ID,
			},
		})
	}
}

// ReplayCommand replays the commands for the ToDo entity
func (t *ToDo) ReplayCommand(cmd entity.CommandRecord) {
	switch cmd.CommandName {
	case "Initialize":
		// Ensure that the correct arguments are passed and task is set
		fmt.Println("Replaying Initialize command", cmd)
		t.Initialize(cmd.Args[0].(string), cmd.Args[1].(string), cmd.Args[2].(string))
		fmt.Println("After init", t.Task)
	case "Complete":
		t.Complete()
	default:
		// Handle unknown commands if necessary
	}
}
