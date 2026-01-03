package work

import "context"

// BusPumperContract defines the interface for pumping work items from a bus.
//
// Implementing this interface is NOT a requirement for having a work bus, but it provides a generic way to
// encapsulate any pumping functionality.
type BusPumperContract interface {
	// Pump pumps work items from the bus for processing.
	Pump(ctx context.Context) error
}

// BusPublisherContract defines the interface for publishing work items to a bus.
type BusPublisherContract interface {
	// Publish publishes a work item to the bus.
	Publish(workItem WorkContract) error
}

// BusSubscriberContract defines the interface for receiving work items from a bus.
type BusSubscriberContract interface {
	// Subscribe receives a work item from the bus.
	//
	// This is generally where the work-specific handlers would be invoked.
	Subscribe(workItem WorkContract) error
}

// BusContract defines the interface for a work bus that can publish and subscribe to work items.
type BusContract interface {
	BusPublisherContract
	BusSubscriberContract
}

// WorkContract defines the interface for a general work item that can be processed.
//
// The intent is for a work item to be self-contained and capable of executing its own processing logic.
type WorkContract interface {
	// Process invokes the functionality to process the work item.
	Process() WorkResultContract

	// Type returns the type or category of the work item.
	Type() WorkType
}

// WorkResultErrorContract defines the interface for retrieving error information from a work result.
type WorkResultErrorContract interface {
	// Error retrieves the error message encountered during work processing. This also allows the implementation to be
	// used as an error type.
	Error() string

	// ErrorInstance retrieves any error encountered during work processing.
	ErrorInstance() error
}

// WorkResultReturnContract defines the interface for retrieving return data from a work result.
type WorkResultReturnContract interface {
	// Return retrieves any relevant data returned from processing the work item.
	Return() any
}

// WorkResultSourceContract defines the interface for retrieving the source work item from a work result.
type WorkResultSourceContract interface {
	// Source returns the source work item associated with this result.
	Source() WorkContract
}

// WorkResultSuccessContract defines the interface for retrieving success status from a work result.
type WorkResultSuccessContract interface {
	// Success indicates whether the work item was processed successfully.
	Success() bool
}

// WorkResultContract defines the interface for the result of processing a work item.
type WorkResultContract interface {
	WorkResultErrorContract
	WorkResultReturnContract
	WorkResultSourceContract
	WorkResultSuccessContract
}
