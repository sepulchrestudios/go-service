package mail

// MessageMetadata holds metadata about a mail message.
type MessageMetadata struct {
	additionalHeaders MessageHeaders
	from              string
	to                string
}

// From returns the sender's mail address.
func (m *MessageMetadata) From() string {
	if m == nil {
		return ""
	}
	return m.from
}

// To returns the recipient's mail address.
func (m *MessageMetadata) To() string {
	if m == nil {
		return ""
	}
	return m.to
}

// AdditionalHeaders returns any additional headers for the mail message.
func (m *MessageMetadata) AdditionalHeaders() MessageHeaders {
	if m == nil || m.additionalHeaders == nil {
		return MessageHeaders{}
	}
	return m.additionalHeaders
}

// NewMessageMetadata creates a new MessageMetadata instance.
func NewMessageMetadata(from string, to string, additionalHeaders MessageHeaders) *MessageMetadata {
	if additionalHeaders == nil {
		additionalHeaders = MessageHeaders{}
	}
	return &MessageMetadata{
		from:              from,
		to:                to,
		additionalHeaders: additionalHeaders,
	}
}
