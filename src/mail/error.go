package mail

import "errors"

// ErrFailedToSendMessage indicates that the mail message failed to send.
var ErrFailedToSendMessage = errors.New("failed to send mail message")
