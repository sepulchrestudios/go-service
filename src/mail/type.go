package mail

import "github.com/sepulchrestudios/go-service/src/work"

// MessageType represents the type or category of a mail message.
type MessageType work.WorkType

const (
	// MessageTypeAll represents all (or no specific) message type(s).
	//
	// This is useful for subscribing to all messages but may not be appropriate when publishing them.
	MessageTypeAll MessageType = "all"
)
