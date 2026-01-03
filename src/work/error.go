package work

import "errors"

// ErrBusCannotBeNil is a sentinel error representing an attempt to use a nil work bus.
var ErrBusCannotBeNil = errors.New("work bus instance cannot be nil")

// ErrCannotProcessWork is a sentinel error representing a work processing failure.
var ErrCannotProcessWork = errors.New("cannot process work")

// ErrCannotPublishWork is a sentinel error representing a work publishing failure.
var ErrCannotPublishWork = errors.New("cannot publish work")

// ErrCannotPumpWork is a sentinel error representing a failure to pump work items for processing.
var ErrCannotPumpWork = errors.New("cannot pump work for processing")

// ErrCannotRegisterWorkHandler is a sentinel error representing a work handler registration failure.
var ErrCannotRegisterWorkHandler = errors.New("cannot register work handler")

// ErrCannotRegisterWorkHandlerNil is a sentinel error representing an attempt to register a nil work handler.
var ErrCannotRegisterWorkHandlerNil = errors.New("cannot register nil work handler")

// ErrCannotSubscribeToWork is a sentinel error representing a work subscription failure.
var ErrCannotSubscribeToWork = errors.New("cannot subscribe to work")
