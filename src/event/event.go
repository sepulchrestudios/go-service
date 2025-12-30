package event

// BusContract defines the interface for an event bus that can publish and subscribe to events.
type BusContract interface {
	// Publish publishes an event to the bus.
	Publish(event EventContract) error

	// Subscribe subscribes to events of a specific type with the provided handler function.
	Subscribe(eventType EventType, handler func(event EventContract) error) error
}

// EventContract defines the interface for a general event that can be processed.
//
// The intent is for an event to be self-contained and capable of executing its own processing logic.
type EventContract interface {
	// Process invokes the functionality to process the event.
	Process() error
}
