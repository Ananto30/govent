package govent

// Event can take all arguments of an event
type Event interface{}

// EventType is event type for an event :3
type EventType int

// EventObject contains event and event type
type EventObject struct {
	EventType
	Event
}

// EventHandler is handler for events and takes any arguments
type EventHandler func(args Event)

// register event types, this is just a format, don't register here,
// rather use your own types like in the example/events.go
const (
	NoobEvent EventType = iota
)
