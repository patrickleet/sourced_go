package entity

import (
	"sourced_go/events"
	"time"
)

// CommandRecord holds information about a digested command
type CommandRecord struct {
	CommandName string
	Args        []interface{}
}

// Event interface defines an event type for the system
type Event interface {
	EventType() string
	GetData() interface{}
}

// EventType returns the event type
func (e LocalEvent) EventType() string {
	return e.Type
}

// GetData returns the event data
func (e LocalEvent) GetData() interface{} {
	return e.Data
}

type LocalEvent struct {
	Type string      // The type of the event (e.g., "Initialized", "Completed")
	Data interface{} // The event's associated data
}

// Entity struct defines the base entity that all domain models will extend
type Entity struct {
	ID              string
	Version         int
	Commands        []CommandRecord
	EventsToEmit    []Event
	Replaying       bool
	SnapshotVersion int
	Timestamp       time.Time
	EventEmitter    *events.EventEmitter // Embed the EventEmitter
}

// NewEntity creates a new base entity with an event emitter
func NewEntity() *Entity {
	return &Entity{
		EventEmitter: events.NewEventEmitter(), // Initialize the embedded EventEmitter
	}
}

// Digest adds a command to the entity if we are not replaying
func (e *Entity) Digest(name string, args ...interface{}) {
	// Only digest new commands if we are not replaying commands
	if e.Replaying {
		return
	}
	e.Commands = append(e.Commands, CommandRecord{
		CommandName: name,
		Args:        args,
	})
	e.Version++
	e.Timestamp = time.Now()
}

func (e *Entity) Enqueue(eventType string, data interface{}) {
	if e.Replaying {
		return
	}
	e.EventsToEmit = append(e.EventsToEmit, LocalEvent{
		Type: eventType,
		Data: data,
	})
}

// EmitQueuedEvents triggers the emission of all queued events
func (e *Entity) EmitQueuedEvents() {
	for _, event := range e.EventsToEmit {
		e.Emit(event.EventType(), event.GetData())
	}
	e.EventsToEmit = nil // Clear the events after emitting
}

// Rehydrate replays the commands to rebuild the entity's state
func (e *Entity) Rehydrate() {
	e.Replaying = true
	for _, commandRecord := range e.Commands {
		e.replayCommand(commandRecord)
	}
	e.Replaying = false
}

// Replay commands (override this in domain models)
func (e *Entity) replayCommand(commandRecord CommandRecord) {
	// Domain-specific models will override this method
}

// EmitNow allows the model to emit events immediately (forwarding call to EventEmitter)
func (e *Entity) Emit(event string, data interface{}) {
	e.EventEmitter.Emit(event, data)
}

// On registers an event listener (forwarding call to EventEmitter)
func (e *Entity) On(event string, listener func(data interface{})) {
	e.EventEmitter.On(event, listener)
}
