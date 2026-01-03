package work

import (
	"context"
	"sync"
)

// Bus is a simple concurrent in-memory implementation of a work bus. It also contains a mutex so it should ONLY be
// passed around by-reference and never by-value.
//
// Under the hood, it essentially implements the Publisher-Subscriber and Observer design patterns.
type Bus struct {
	handlers   map[WorkType][]HandlerFunc
	handlersMu sync.Mutex
	pipeline   chan WorkContract
	resultChan chan WorkResultContract
}

// Publish publishes a work item to the bus.
func (b *Bus) Publish(workItem WorkContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.pipeline == nil {
		return ErrCannotPublishWork
	}
	if workItem == nil {
		return nil
	}
	b.pipeline <- workItem
	return nil
}

// Pump continuously pumps work from the internal pipeline for processing until the provided context is done.
//
// This method BLOCKS until ctx.Done() is closed, so it should be run in its own goroutine.
func (b *Bus) Pump(ctx context.Context) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.pipeline == nil {
		return ErrCannotPumpWork
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case workItem := <-b.pipeline:
			if b.resultChan == nil {
				b.resultChan = make(chan WorkResultContract)
			}
			if workItem != nil {
				// Process each work item in its own goroutine to avoid creating a blocking queue.
				go func(w WorkContract) {
					results := b.Subscribe(w)
					for _, result := range results {
						b.resultChan <- result
					}
				}(workItem)
			}
		}
	}
}

// Subscribe receives a work item from the bus and processes it.
func (b *Bus) Subscribe(workItem WorkContract) []WorkResultContract {
	if b == nil || b.handlers == nil || workItem == nil {
		return []WorkResultContract{}
	}

	// Ensure we don't get a collision if two or more goroutines try to read concurrently
	b.handlersMu.Lock()
	defer b.handlersMu.Unlock()

	// Set up a consistent way to process our work handlers.
	processingResultChan := make(chan WorkResultContract)
	processHandlersFunc := func(handlers []HandlerFunc) {
		for _, handler := range handlers {
			if handler != nil {
				// Process each work handler in its own goroutine to avoid creating a blocking queue.
				go func(handlerFunc HandlerFunc) {
					result := handlerFunc(workItem)
					if result != nil {
						processingResultChan <- result
					}
				}(handler)
			}
		}
	}

	// Invoke handlers registered for the specific work type first.
	if handlers, exists := b.handlers[workItem.Type()]; exists {
		processHandlersFunc(handlers)
	}
	// Invoke any handlers registered for "all" work types second.
	if handlers, exists := b.handlers[WorkTypeAll]; exists {
		processHandlersFunc(handlers)
	}

	// Return all errors (if any) encountered during work processing.
	close(processingResultChan)
	results := []WorkResultContract{}
	for result := range processingResultChan {
		results = append(results, result)
	}
	return results
}

// RegisterHandler registers a handler function for a specific work type.
func (b *Bus) RegisterHandler(workType WorkType, handler HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if handler == nil {
		return ErrCannotRegisterWorkHandlerNil
	}
	if b.handlers == nil {
		b.handlers = make(map[WorkType][]HandlerFunc)
	}
	if _, exists := b.handlers[workType]; !exists {
		b.handlers[workType] = []HandlerFunc{}
	}
	b.handlers[workType] = append(b.handlers[workType], handler)
	return nil
}

// Results returns a channel that emits results from work processing.
func (b *Bus) Results() chan WorkResultContract {
	if b == nil {
		return nil
	}
	return b.resultChan
}

// NewBus creates a new concurrent work bus instance.
func NewBus() *Bus {
	return &Bus{
		handlers:   make(map[WorkType][]HandlerFunc),
		pipeline:   make(chan WorkContract),
		resultChan: make(chan WorkResultContract),
	}
}
