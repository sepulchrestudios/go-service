package mail

import (
	"context"

	"github.com/sepulchrestudios/go-service/src/work"
)

// Bus is a simple in-memory implementation of a mail bus.
//
// Under the hood, it essentially implements the Publisher-Subscriber and Observer design patterns.
type Bus struct {
	workBus work.PumpingWorkHandlerBusContract
}

// Publish publishes a mail message to the bus.
func (b *Bus) Publish(message work.WorkContract) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	if message == nil {
		return nil
	}
	return b.workBus.Publish(message)
}

// Pump continuously pumps mail messages from the internal pipeline for processing.
func (b *Bus) Pump(ctx context.Context) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	return b.workBus.Pump(ctx)
}

// Subscribe receives a mail message from the bus and processes it.
func (b *Bus) Subscribe(message work.WorkContract) []work.WorkResultContract {
	if b == nil || b.workBus == nil || message == nil {
		return []work.WorkResultContract{}
	}
	return b.workBus.Subscribe(message)
}

// RegisterDefaultHandler registers a default handler function for ALL mail message types.
func (b *Bus) RegisterDefaultHandler() error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	return b.RegisterMessageHandler(MessageTypeAll, func(message work.WorkContract) work.WorkResultContract {
		if message == nil {
			return nil
		}
		return message.Process()
	})
}

// RegisterMessageHandler registers a handler function for a specific mail message type.
func (b *Bus) RegisterMessageHandler(messageType MessageType, handler work.HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	return b.RegisterHandler(work.WorkType(messageType), handler)
}

// RegisterHandler registers a handler function for a specific mail message type.
func (b *Bus) RegisterHandler(messageType work.WorkType, handler work.HandlerFunc) error {
	if b == nil {
		return ErrBusCannotBeNil
	}
	if b.workBus == nil {
		return ErrWorkBusCannotBeNil
	}
	if handler == nil {
		return ErrCannotRegisterMessageHandlerNil
	}
	return b.workBus.RegisterHandler(messageType, handler)
}

// Results returns a channel that emits results from mail message processing.
func (b *Bus) Results() chan work.WorkResultContract {
	if b == nil {
		return nil
	}
	if b.workBus == nil {
		return nil
	}
	return b.workBus.Results()
}

// NewBus creates a new mail bus instance.
func NewBus(workBus work.PumpingWorkHandlerBusContract) *Bus {
	return &Bus{
		workBus: workBus,
	}
}
