package mail

import "github.com/sepulchrestudios/go-service/src/work"

// MessageContentBodyContract defines the contract for representing message body content.
type MessageContentBodyContract interface {
	// Body returns the body of the message.
	Body() []byte
}

// MessageContentSubjectContract defines the contract for representing message subject content.
type MessageContentSubjectContract interface {
	// Subject returns the subject of the message.
	Subject() string
}

// MessageContentContract defines the contract for representing message content.
type MessageContentContract interface {
	MessageContentBodyContract
	MessageContentSubjectContract
}

// MessageMetadataHeadersContract defines the contract for representing message headers.
type MessageMetadataHeadersContract interface {
	// Headers returns the headers as a map of string key-value pairs.
	Headers() MessageHeaders
}

// MessageMetadataRecipientContract defines the contract for representing message recipient information.
type MessageMetadataRecipientContract interface {
	// To returns the recipient's address.
	To() string
}

type MessageMetadataSenderContract interface {
	// From returns the sender's address.
	From() string
}

// MessageMetadataContract defines the contract for representing message metadata.
type MessageMetadataContract interface {
	MessageMetadataRecipientContract
	MessageMetadataSenderContract
}

// MessageMetadataWithHeadersContract defines the contract for representing message metadata with headers.
type MessageMetadataWithHeadersContract interface {
	MessageMetadataContract
	MessageMetadataHeadersContract
}

// MailMessageContentContract defines the contract for representing mail message content.
type MailMessageContentContract interface {
	// Content returns the content of the mail message.
	Content() MessageContentContract
}

// MailMessageMetadataContract defines the contract for representing mail message metadata.
type MailMessageMetadataContract interface {
	// Metadata returns the metadata of the mail message.
	Metadata() MessageMetadataWithHeadersContract
}

// MailMessageContract defines the contract for representing a mail message.
type MailMessageContract interface {
	MailMessageContentContract
	MailMessageMetadataContract
}

// MailerContract defines the contract for sending mail messages.
type MailerContract interface {
	// Send sends the given mail message(s) and returns the matching result(s) from the operation.
	Send(messages []MailMessageContract) ([]work.WorkResultContract, error)
}
