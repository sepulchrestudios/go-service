package event

import (
	"context"

	"github.com/sepulchrestudios/go-service/src/work"
)

// Bus is a simple in-memory implementation of an event bus.
//
// Under the hood, it essentially implements the Publisher-Subscriber and Observer design patterns.
type Bus struct {
	workBus work.PumpingWorkHandlerBusContract
}

// Publish publishes an event to the bus.
func (b *Bus) Publish(event work.WorkContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	if event == nil {
		return nil
	}
	return b.workBus.Publish(event)
}

// Pump continuously pumps events from the internal pipeline for processing.
func (b *Bus) Pump(ctx context.Context) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	return b.workBus.Pump(ctx)
}

// Subscribe receives an event from the bus and processes it.
func (b *Bus) Subscribe(event work.WorkContract) []work.WorkResultContract {
	if b == nil || b.workBus == nil || event == nil {
		return []work.WorkResultContract{}
	}
	return b.workBus.Subscribe(event)
}

// RegisterDefaultHandler registers a default handler function for ALL event types.
func (b *Bus) RegisterDefaultHandler() error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	return b.RegisterEventHandler(EventTypeAll, func(event work.WorkContract) work.WorkResultContract {
		if event == nil {
			return nil
		}
		return event.Process()
	})
}

// RegisterEventHandler registers a handler function for a specific event type.
func (b *Bus) RegisterEventHandler(eventType EventType, handler work.HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	return b.RegisterHandler(work.WorkType(eventType), handler)
}

// RegisterHandler registers a handler function for a specific work type.
func (b *Bus) RegisterHandler(eventType work.WorkType, handler work.HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	if handler == nil {
		return ErrCannotRegisterEventHandlerNil
	}
	return b.workBus.RegisterHandler(eventType, handler)
}

// Results returns a channel that emits results from event processing.
func (b *Bus) Results() chan work.WorkResultContract {
	if b == nil {
		return nil
	}
	if b.workBus == nil {
		return nil
	}
	return b.workBus.Results()
}

// NewBus creates a new event bus instance.
func NewBus(workBus work.PumpingWorkHandlerBusContract) *Bus {
	return &Bus{
		workBus: workBus,
	}
}
