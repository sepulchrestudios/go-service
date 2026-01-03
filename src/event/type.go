package event

import "github.com/sepulchrestudios/go-service/src/work"

// EventType represents the type or category of an event.
type EventType work.WorkType

const (
	// EventTypeAll represents all (or no specific) event type(s).
	//
	// This is useful for subscribing to all events but may not be appropriate when publishing them.
	EventTypeAll EventType = "all"
)
