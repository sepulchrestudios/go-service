package mail

// MessageContent holds the content of an mail message.
type MessageContent struct {
	body    []byte
	subject string
}

// Body returns the body of the mail message.
func (c *MessageContent) Body() []byte {
	if c == nil {
		return []byte{}
	}
	return c.body
}

// Subject returns the subject of the mail message.
func (c *MessageContent) Subject() string {
	if c == nil {
		return ""
	}
	return c.subject
}

// NewMessageContent creates a new MessageContent instance.
func NewMessageContent(subject string, body []byte) *MessageContent {
	if body == nil {
		body = []byte{}
	}
	return &MessageContent{
		subject: subject,
		body:    body,
	}
}
