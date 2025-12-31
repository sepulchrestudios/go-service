package event

import (
	"context"
	"fmt"
	"sync"
)

// HandlerFunc defines the function signature for event handler functions.
type HandlerFunc func(event EventContract) error

// Bus is a simple in-memory implementation of an event bus. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
type Bus struct {
	handlers   map[EventType][]HandlerFunc
	handlersMu sync.Mutex
	pipeline   chan EventContract
	errorChan  chan error
}

// Errors returns a channel that emits errors encountered during event processing.
func (b *Bus) Errors() chan error {
	if b == nil {
		return nil
	}
	return b.errorChan
}

// Publish publishes an event to the bus.
func (b *Bus) Publish(event EventContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.pipeline == nil {
		return ErrCannotPublishEvent
	}
	if event == nil {
		return nil
	}
	b.pipeline <- event
	return nil
}

// PumpEvents continuously pumps events from the internal pipeline for processing until the provided context is done.
//
// This method BLOCKS until ctx.Done() is closed, so it should be run in its own goroutine.
func (b *Bus) PumpEvents(ctx context.Context) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.pipeline == nil {
		return ErrCannotPumpEvents
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-b.pipeline:
			if b.errorChan == nil {
				b.errorChan = make(chan error)
			}
			if event != nil {
				// Process each event in its own goroutine to avoid creating a blocking queue.
				go func(e EventContract) {
					if err := b.Receive(e); err != nil {
						b.errorChan <- err
					}
				}(event)
			}
		}
	}
}

// Receive receives an event from the bus and processes it.
func (b *Bus) Receive(event EventContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.handlers == nil || event == nil {
		return nil
	}

	// Ensure we don't get a collision if two or more goroutines try to read concurrently
	b.handlersMu.Lock()
	defer b.handlersMu.Unlock()

	// Set up a consistent way to process our event handlers.
	processingErrChan := make(chan error)
	processHandlersFunc := func(handlers []HandlerFunc) {
		for _, handler := range handlers {
			if handler != nil {
				// Process each event handler in its own goroutine to avoid creating a blocking queue.
				go func(handlerFunc HandlerFunc) {
					if err := handlerFunc(event); err != nil {
						processingErrChan <- err
					}
				}(handler)
			}
		}
	}

	// Invoke handlers registered for the specific event type first.
	if handlers, exists := b.handlers[event.Type()]; exists {
		processHandlersFunc(handlers)
	}
	// Invoke any handlers registered for "all" event types second.
	if handlers, exists := b.handlers[EventTypeAll]; exists {
		processHandlersFunc(handlers)
	}

	// Return all errors (if any) encountered during event processing.
	close(processingErrChan)
	var fullErr error
	for err := range processingErrChan {
		if fullErr == nil {
			fullErr = err
		} else {
			fullErr = fmt.Errorf("%w: %w", fullErr, err)
		}
	}
	return fullErr
}

// RegisterDefaultHandler registers a default handler function for ALL event types.
func (b *Bus) RegisterDefaultHandler() error {
	return b.RegisterHandler(EventTypeAll, func(event EventContract) error {
		if event == nil {
			return nil
		}
		return event.Process()
	})
}

// RegisterHandler registers a handler function for a specific event type.
func (b *Bus) RegisterHandler(eventType EventType, handler HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if handler == nil {
		return ErrCannotRegisterEventHandlerNil
	}
	if b.handlers == nil {
		b.handlers = make(map[EventType][]HandlerFunc)
	}
	if _, exists := b.handlers[eventType]; !exists {
		b.handlers[eventType] = []HandlerFunc{}
	}
	b.handlers[eventType] = append(b.handlers[eventType], handler)
	return nil
}

// NewBus creates a new event bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers:  make(map[EventType][]HandlerFunc),
		pipeline:  make(chan EventContract),
		errorChan: make(chan error),
	}
}
