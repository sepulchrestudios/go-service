package mail

import "github.com/sepulchrestudios/go-service/src/work"

// MessageContentContract defines the contract for representing mail message content.
type MessageContentContract interface {
	// Body returns the body of the mail message.
	Body() []byte

	// Subject returns the subject of the mail message.
	Subject() string
}

// MessageMetadataContract defines the contract for representing mail message metadata.
type MessageMetadataContract interface {
	// AdditionalHeaders returns any additional headers for the mail message.
	AdditionalHeaders() MessageHeaders

	// From returns the sender's mail address.
	From() string

	// To returns the recipient's mail address.
	To() string
}

// MessageContract defines the contract for representing a mail message.
type MessageContract interface {
	// Content returns the content of the mail message.
	Content() MessageContentContract

	// Metadata returns the metadata of the mail message.
	Metadata() MessageMetadataContract
}

// MailerContract defines the contract for sending mail messages.
type MailerContract interface {
	// Send sends the given mail message(s) and returns the matching result(s) from the operation.
	Send(messages []MessageContract) ([]work.WorkResultContract, error)
}
