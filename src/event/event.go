package event

// BusPublisherContract defines the interface for publishing events to a bus.
type BusPublisherContract interface {
	// Publish publishes an event to the bus.
	Publish(event EventContract) error
}

// BusReciverContract defines the interface for receiving events from a bus.
type BusReciverContract interface {
	// Receive receives an event from the bus.
	//
	// This is generally where the event-specific handlers would be invoked.
	Receive(event EventContract) error
}

// BusContract defines the interface for an event bus that can publish and subscribe to events.
type BusContract interface {
	BusPublisherContract
	BusReciverContract
}

// EventContract defines the interface for a general event that can be processed.
//
// The intent is for an event to be self-contained and capable of executing its own processing logic.
type EventContract interface {
	// Process invokes the functionality to process the event.
	Process() error

	// Type returns the type or category of the event.
	Type() EventType
}
