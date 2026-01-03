package event

import "errors"

// ErrBusCannotBeNil is a sentinel error representing an attempt to use a nil event bus.
var ErrBusCannotBeNil = errors.New("event bus instance cannot be nil")

// ErrCannotRegisterEventHandlerNil is a sentinel error representing an attempt to register a nil event handler.
var ErrCannotRegisterEventHandlerNil = errors.New("cannot register nil event handler")

// ErrWorkBusCannotBeNil is a sentinel error representing an attempt to use a nil underlying work bus.
var ErrWorkBusCannotBeNil = errors.New("underlying work bus instance cannot be nil")
