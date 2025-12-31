package event

import (
	"context"
	"sync"
)

// HandlerFunc defines the function signature for event handler functions.
type HandlerFunc func(event EventContract) EventResultContract

// Bus is a simple concurrent in-memory implementation of an event bus. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
//
// Under the hood, it essentially implements the Publisher-Subscriber and Observer design patterns.
type Bus struct {
	handlers   map[EventType][]HandlerFunc
	handlersMu sync.Mutex
	pipeline   chan EventContract
	resultChan chan EventResultContract
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
			if b.resultChan == nil {
				b.resultChan = make(chan EventResultContract)
			}
			if event != nil {
				// Process each event in its own goroutine to avoid creating a blocking queue.
				go func(e EventContract) {
					results := b.Subscribe(e)
					for _, result := range results {
						b.resultChan <- result
					}
				}(event)
			}
		}
	}
}

// Subscribe receives an event from the bus and processes it.
func (b *Bus) Subscribe(event EventContract) []EventResultContract {
	if b == nil || b.handlers == nil || event == nil {
		return []EventResultContract{}
	}

	// Ensure we don't get a collision if two or more goroutines try to read concurrently
	b.handlersMu.Lock()
	defer b.handlersMu.Unlock()

	// Set up a consistent way to process our event handlers.
	processingResultChan := make(chan EventResultContract)
	processHandlersFunc := func(handlers []HandlerFunc) {
		for _, handler := range handlers {
			if handler != nil {
				// Process each event handler in its own goroutine to avoid creating a blocking queue.
				go func(handlerFunc HandlerFunc) {
					result := handlerFunc(event)
					if result != nil {
						processingResultChan <- result
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
	close(processingResultChan)
	results := []EventResultContract{}
	for result := range processingResultChan {
		results = append(results, result)
	}
	return results
}

// RegisterDefaultHandler registers a default handler function for ALL event types.
func (b *Bus) RegisterDefaultHandler() error {
	return b.RegisterHandler(EventTypeAll, func(event EventContract) EventResultContract {
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

// Results returns a channel that emits results from event processing.
func (b *Bus) Results() chan EventResultContract {
	if b == nil {
		return nil
	}
	return b.resultChan
}

// NewBus creates a new event bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers:   make(map[EventType][]HandlerFunc),
		pipeline:   make(chan EventContract),
		resultChan: make(chan EventResultContract),
	}
}
