package message

import "github.com/JordanRad/chatbook/services/cmd/chat-service/db/models"

type MessageSocket struct {
	store Store
}

func NewMessageSocket(s Store) *MessageSocket {
	return &MessageSocket{
		store: s,
	}
}

// Compile time assertion that this service implements the generated interface
var _ Service = (*MessageSocket)(nil)

func (ms *MessageSocket) ReceiveMessage(m models.ConversationMessage) error {
	return nil
}
