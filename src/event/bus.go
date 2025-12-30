package event

import (
	"context"
	"fmt"
	"sync"
)

// Bus is a simple in-memory implementation of an event bus. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
type Bus struct {
	handlers   map[EventType][]func(event EventContract) error
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
func (b *Bus) PumpEvents(ctx context.Context) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.pipeline == nil {
		return ErrCannotPumpEvents
	}
	// TODO: use goroutines to make this concurrent
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-b.pipeline:
			if b.errorChan == nil {
				b.errorChan = make(chan error)
			}
			if event != nil {
				if err := b.Receive(event); err != nil {
					b.errorChan <- err
				}
			}
		}
	}
}

// Receive receives an event from the bus and processes it.
func (b *Bus) Receive(event EventContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if event == nil {
		return nil
	}
	if b.handlers == nil {
		b.handlers = make(map[EventType][]func(event EventContract) error)
	}

	// Ensure we don't get a collision if two or more goroutines try to read concurrently
	b.handlersMu.Lock()
	defer b.handlersMu.Unlock()

	// TODO: use goroutines to make this concurrent

	// Invoke handlers registered for the specific event type first.
	processingErrChan := make(chan error)
	if handlers, exists := b.handlers[event.Type()]; exists {
		for _, handler := range handlers {
			if handler != nil {
				if err := handler(event); err != nil {
					processingErrChan <- err
				}
			}
		}
	}
	// Invoke any handlers registered for "all" event types second.
	if handlers, exists := b.handlers[EventTypeAll]; exists {
		for _, handler := range handlers {
			if handler != nil {
				if err := handler(event); err != nil {
					processingErrChan <- err
				}
			}
		}
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

// RegisterHandler registers a handler function for a specific event type.
func (b *Bus) RegisterHandler(eventType EventType, handler func(event EventContract) error) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if handler == nil {
		return ErrCannotRegisterEventHandlerNil
	}
	if b.handlers == nil {
		b.handlers = make(map[EventType][]func(event EventContract) error)
	}
	if _, exists := b.handlers[eventType]; !exists {
		b.handlers[eventType] = []func(event EventContract) error{}
	}
	b.handlers[eventType] = append(b.handlers[eventType], handler)
	return nil
}

// NewBus creates a new event bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers:  make(map[EventType][]func(event EventContract) error),
		pipeline:  make(chan EventContract),
		errorChan: make(chan error),
	}
}
