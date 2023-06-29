package message

import "github.com/JordanRad/chatbook/services/cmd/chat-service/db/models"

// Used to implement for the store in the service
type Store interface {
	AddMessage(m models.ConversationMessage) error
}

type Service interface {
	ReceiveMessage(m models.ConversationMessage) error
}
