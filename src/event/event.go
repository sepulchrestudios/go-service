package event

// BusPublisherContract defines the interface for publishing events to a bus.
type BusPublisherContract interface {
	// Publish publishes an event to the bus.
	Publish(event EventContract) error
}

// BusReceiverContract defines the interface for receiving events from a bus.
type BusReceiverContract interface {
	// Receive receives an event from the bus.
	//
	// This is generally where the event-specific handlers would be invoked.
	Receive(event EventContract) error
}

// BusContract defines the interface for an event bus that can publish and subscribe to events.
type BusContract interface {
	BusPublisherContract
	BusReceiverContract
}

// EventContract defines the interface for a general event that can be processed.
//
// The intent is for an event to be self-contained and capable of executing its own processing logic.
type EventContract interface {
	// Process invokes the functionality to process the event.
	Process() EventResultContract

	// Type returns the type or category of the event.
	Type() EventType
}

// EventResultErrorContract defines the interface for retrieving error information from an event result.
type EventResultErrorContract interface {
	// Error retrieves the error message encountered during event processing. This also allows the implementation to be
	// used as an error type.
	Error() string

	// ErrorInstance retrieves any error encountered during event processing.
	ErrorInstance() error
}

// EventResultReturnContract defines the interface for retrieving return data from an event result.
type EventResultReturnContract interface {
	// Return retrieves any relevant data returned from processing the event.
	Return() any
}

// EventResultSuccessContract defines the interface for retrieving success status from an event result.
type EventResultSuccessContract interface {
	// Success indicates whether the event was processed successfully.
	Success() bool
}

// EventResultContract defines the interface for the result of processing an event.
type EventResultContract interface {
	EventResultErrorContract
	EventResultReturnContract
	EventResultSuccessContract
}
