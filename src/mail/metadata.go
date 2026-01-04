package mail

// MessageMetadata holds metadata about a mail message.
type MessageMetadata struct {
	headers MessageHeaders
	from    string
	to      string
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

// Headers returns any headers for the mail message.
func (m *MessageMetadata) Headers() MessageHeaders {
	if m == nil || m.headers == nil {
		return MessageHeaders{}
	}
	return m.headers
}

// NewMessageMetadata creates a new MessageMetadata instance.
func NewMessageMetadata(from string, to string, headers MessageHeaders) *MessageMetadata {
	if headers == nil {
		headers = MessageHeaders{}
	}
	return &MessageMetadata{
		from:    from,
		to:      to,
		headers: headers,
	}
}
