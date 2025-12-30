package event

import "errors"

// ErrCannotProcessEvent is a sentinel error representing an event processing failure.
var ErrCannotProcessEvent = errors.New("cannot process event")

// ErrCannotPublishEvent is a sentinel error representing an event publishing failure.
var ErrCannotPublishEvent = errors.New("cannot publish event")

// ErrCannotSubscribeToEvent is a sentinel error representing an event subscription failure.
var ErrCannotSubscribeToEvent = errors.New("cannot subscribe to event")
