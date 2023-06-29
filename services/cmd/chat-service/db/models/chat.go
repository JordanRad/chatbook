package models

import "time"

type ConversationMessage struct {
	SenderID string    `json:"senderID"`
	TS       time.Time `json:"timestamp"`
	Content  string    `json:"content"`
}

type Conversation struct {
	ID, LastMessageSenderID, LastMessageContent, LastMessageDeliveredAt, OtherParticipantID string
}
