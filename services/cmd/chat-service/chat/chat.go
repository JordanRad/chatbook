package chat

import (
	"context"
	"log"

	"github.com/JordanRad/chatbook/services/internal/gen/chat"
)

type Service struct {
	logger *log.Logger
	store  Store
}

type Store interface {
	GetHistory(ctx context.Context) error
}

// Compile time assertion that this service implements the generated interface
var _ chat.Service = (*Service)(nil)

func NewService(logger *log.Logger, store Store) *Service {
	return &Service{
		logger: logger,
		store:  store,
	}
}

func (s *Service) GetChatHistory(ctx context.Context, p *chat.GetChatHistoryPayload) (*chat.ChatHistoryResponse, error) {

	response := &chat.ChatHistoryResponse{
		ID: "111",
	}
	return response, nil
}
