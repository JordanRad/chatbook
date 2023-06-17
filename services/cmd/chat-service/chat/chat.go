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

func (s *Service) GetConversationHistory(ctx context.Context, p *chat.GetConversationHistoryPayload) (*chat.ChatHistoryResponse, error) {
	response := &chat.ChatHistoryResponse{
		ID: "111",
	}
	return response, nil
}

func (s *Service) SearchInConversation(ctx context.Context, p *chat.SearchInConversationPayload) (*chat.ChatHistoryResponse, error) {
	return nil, nil
}

func (s *Service) GetConversationsList(ctx context.Context, p *chat.GetConversationsListPayload) (*chat.ConversationsListResponse, error) {
	return nil, nil
}

func (s *Service) AddConversation(ctx context.Context, p *chat.AddConversationPayload) (*chat.OperationStatusResponse, error) {
	return nil, nil
}
