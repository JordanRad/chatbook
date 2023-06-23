package models

type ConversationMessage struct {
	SenderID string `json:"senderID"`
	TS       int64  `json:"timestamp"`
	Content  string `json:"content"`
}

type Conversation struct {
	ID, LastMessageSenderID, LastMessageContent, LastMessageDeliveredAt, OtherParticipantID string
}
