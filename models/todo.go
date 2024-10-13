package models

type ToDoInitialized struct {
	ID      string
	Address string
	Task    string
}

func (e ToDoInitialized) EventType() string {
	return "ToDoInitialized"
}

type ToDoCompleted struct {
	ID string
}

func (e ToDoCompleted) EventType() string {
	return "ToDoCompleted"
}

type ToDo struct {
	*Entity
	Address   string
	Task      string
	Completed bool
	Removed   bool
}

func NewToDo() *ToDo {
	return &ToDo{
		Entity: NewEntity(),
	}
}

func (t *ToDo) Initialize(id, address, task string) {
	t.ID = id
	t.Address = address
	t.Task = task
	t.Completed = false
	t.Removed = false
	t.DigestCommand("Initialize", id, address, task)
	t.Enqueue(ToDoInitialized{id, address, task})
}

func (t *ToDo) Complete() {
	if !t.Completed {
		t.Completed = true
		t.DigestCommand("Complete", t.ID)
		t.Enqueue(ToDoCompleted{t.ID})
	}
}
