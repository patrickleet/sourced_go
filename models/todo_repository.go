package models

// ToDoRepository inherits from Repository
type ToDoRepository struct {
	*Repository // Inherits EventEmitter and other repository methods
}

// NewToDoRepository initializes a new ToDoRepository
func NewToDoRepository() *ToDoRepository {
	return &ToDoRepository{
		Repository: NewRepository(), // Initialize the base repository
	}
}

// SaveToDo saves the ToDo entity and commits events
func (r *ToDoRepository) SaveToDo(todo *ToDo) {
	r.Save(todo.Entity)
}
