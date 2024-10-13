package events

import "sync"

// EventEmitter struct for event subscription and emission
type EventEmitter struct {
	listeners map[string][]func(data interface{})
	mu        sync.RWMutex
}

// NewEventEmitter initializes a new event emitter
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]func(data interface{})),
	}
}

// On registers an event listener for a given event
func (emitter *EventEmitter) On(event string, listener func(data interface{})) {
	emitter.mu.Lock()
	defer emitter.mu.Unlock()
	emitter.listeners[event] = append(emitter.listeners[event], listener)
}

// Emit triggers all listeners for an event
func (emitter *EventEmitter) Emit(event string, data interface{}) {
	emitter.mu.RLock()
	defer emitter.mu.RUnlock()
	if listeners, found := emitter.listeners[event]; found {
		for _, listener := range listeners {
			listener(data)
		}
	}
}

// RemoveListener removes a listener for an event
func (emitter *EventEmitter) RemoveListener(event string, listenerToRemove func(data interface{})) {
	emitter.mu.Lock()
	defer emitter.mu.Unlock()
	if listeners, found := emitter.listeners[event]; found {
		for i, listener := range listeners {
			if &listener == &listenerToRemove {
				emitter.listeners[event] = append(listeners[:i], listeners[i+1:]...)
				break
			}
		}
	}
}
