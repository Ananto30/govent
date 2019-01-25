package govent

import "fmt"

var eventMap map[EventType][]EventHandler

func init() {
	eventMap = make(map[EventType][]EventHandler)
}

// Subscribe to a function
func Subscribe(eH EventHandler, eT EventType) {
	if len(eventMap[eT]) == 0 {
		eventMap[eT] = make([]EventHandler, 0)
	}
	eventMap[eT] = append(eventMap[eT], eH)
}

// Publish event
func Publish(eT EventType, e Event) {
	for _, eH := range eventMap[eT] {
		eH(e)
	}
}

// Run is a goroutine for receving and publishing events
func Run(publisher chan EventObject) {
	for {
		eventObject := <-publisher
		fmt.Println("Event received ", eventObject.EventType)
		Publish(eventObject.EventType, eventObject.Event)
	}
}

// NewEventPublisher makes a new publisher channel for events
func NewEventPublisher() chan EventObject {
	publisher := make(chan EventObject)
	go Run(publisher)
	return publisher
}
