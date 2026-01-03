package mail

import "errors"

// ErrBusCannotBeNil is a sentinel error representing an attempt to use a nil mail bus.
var ErrBusCannotBeNil = errors.New("mail bus instance cannot be nil")

// ErrCannotRegisterMessageHandlerNil is a sentinel error representing an attempt to register a nil message handler.
var ErrCannotRegisterMessageHandlerNil = errors.New("cannot register nil message handler")

// ErrFailedToSendMessage indicates that the mail message failed to send.
var ErrFailedToSendMessage = errors.New("failed to send mail message")

// ErrWorkBusCannotBeNil is a sentinel error representing an attempt to use a nil underlying work bus.
var ErrWorkBusCannotBeNil = errors.New("underlying work bus instance cannot be nil")
