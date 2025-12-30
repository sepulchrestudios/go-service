package event

import "errors"

// ErrBusCannotBeNil is a sentinel error representing an attempt to use a nil event bus.
var ErrBusCannotBeNil = errors.New("event bus instance cannot be nil")

// ErrCannotProcessEvent is a sentinel error representing an event processing failure.
var ErrCannotProcessEvent = errors.New("cannot process event")

// ErrCannotPublishEvent is a sentinel error representing an event publishing failure.
var ErrCannotPublishEvent = errors.New("cannot publish event")

// ErrCannotPumpEvents is a sentinel error representing a failure to pump events for processing.
var ErrCannotPumpEvents = errors.New("cannot pump events for processing")

// ErrCannotRegisterEventHandler is a sentinel error representing an event handler registration failure.
var ErrCannotRegisterEventHandler = errors.New("cannot register event handler")

// ErrCannotRegisterEventHandlerNil is a sentinel error representing an attempt to register a nil event handler.
var ErrCannotRegisterEventHandlerNil = errors.New("cannot register nil event handler")

// ErrCannotSubscribeToEvent is a sentinel error representing an event subscription failure.
var ErrCannotSubscribeToEvent = errors.New("cannot subscribe to event")
