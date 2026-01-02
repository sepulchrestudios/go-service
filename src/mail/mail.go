package mail

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

// MailerResultErrorContract defines the interface for retrieving error information from a mail sending result.
type MailerResultErrorContract interface {
	// Error retrieves the error message encountered during mail sending. This also allows the implementation to be
	// used as an error type.
	Error() string

	// ErrorInstance retrieves any error encountered during mail sending.
	ErrorInstance() error
}

// MailerResultReturnContract defines the interface for retrieving return data from a mail sending result.
type MailerResultReturnContract interface {
	// Return retrieves any relevant data returned from processing the mail sending operation.
	Return() any
}

// MailerResultSourceContract defines the interface for retrieving the source message from a mail sending result.
type MailerResultSourceContract interface {
	// Source returns the source message associated with this result.
	Source() MessageContract
}

// MailerResultSuccessContract defines the interface for retrieving success status from a mail sending result.
type MailerResultSuccessContract interface {
	// Success indicates whether the mail sending operation was successful.
	Success() bool
}

// MailerResultContract defines the contract for representing the result of a mail sending operation.
type MailerResultContract interface {
	MailerResultErrorContract
	MailerResultReturnContract
	MailerResultSourceContract
	MailerResultSuccessContract
}

// MailerContract defines the contract for sending mail messages.
type MailerContract interface {
	// Send sends the given mail message(s) and returns the matching result(s) from the operation.
	Send(messages []MessageContract) ([]MailerResultContract, error)
}
