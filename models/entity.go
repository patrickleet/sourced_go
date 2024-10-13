package models

import "time"

type CommandRecord struct {
	CommandName string
	Args        []interface{}
}

type Event interface {
	EventType() string
}

type Entity struct {
	ID              string
	Version         int
	Commands        []CommandRecord
	EventsToEmit    []Event
	Replaying       bool
	SnapshotVersion int
	Timestamp       time.Time
	EventEmitter    *EventEmitter
}

func NewEntity() *Entity {
	return &Entity{
		EventEmitter: NewEventEmitter(),
	}
}

func (e *Entity) DigestCommand(commandName string, args ...interface{}) {
	e.Commands = append(e.Commands, CommandRecord{
		CommandName: commandName,
		Args:        args,
	})
	e.Version++
	e.Timestamp = time.Now()
}

func (e *Entity) Enqueue(event Event) {
	e.EventsToEmit = append(e.EventsToEmit, event)
}

func (e *Entity) Commit() {
	e.emitEvents()
	e.EventsToEmit = []Event{}
}

func (e *Entity) emitEvents() {
	for _, event := range e.EventsToEmit {
		e.EventEmitter.Emit(event.EventType(), event)
	}
}

func (e *Entity) Rehydrate() {
	e.Replaying = true
	for _, commandRecord := range e.Commands {
		e.replayCommand(commandRecord)
	}
	e.Replaying = false
}

func (e *Entity) replayCommand(commandRecord CommandRecord) {
	// Override this in derived entities
}
