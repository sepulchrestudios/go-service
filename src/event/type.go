package event

// EventType represents the type or category of an event.
type EventType string

const (
	// EventTypeAll represents all (or no specific) event type(s).
	//
	// This is useful for subscribing to all events but may not be appropriate when publishing them.
	EventTypeAll EventType = "all"
)
