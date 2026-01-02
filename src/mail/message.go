package mail

// Message represents a mail message.
type Message struct {
	content  *MessageContent
	metadata *MessageMetadata
}

// Content returns the content of the mail message.
func (m *Message) Content() *MessageContent {
	if m == nil {
		return nil
	}
	return m.content
}

// Metadata returns the metadata of the mail message.
func (m *Message) Metadata() *MessageMetadata {
	if m == nil {
		return nil
	}
	return m.metadata
}

// NewMessage creates a new mail message with the given metadata and content.
func NewMessage(metadata *MessageMetadata, content *MessageContent) *Message {
	return &Message{
		metadata: metadata,
		content:  content,
	}
}
