package chat

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JordanRad/chatbook/services/cmd/chat-service/db/models"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/gen/chat"
)

type Service struct {
	logger *log.Logger
	store  Store
}

type Store interface {
	FindHistoryByID(ctx context.Context, ID string, limit int, beforeTS string) ([]models.ConversationMessage, error)
	FindBySearchInput(ctx context.Context, ID, input string, limit int) ([]models.ConversationMessage, error)
	ListConversationsByUserID(ctx context.Context, userID string, limit int) ([]models.Conversation, error)
	CreateConversation(ctx context.Context, participants []auth.FriendsList) error
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
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting user from context: %w", err)
	}
	fmt.Println(u)

	ts := time.Unix(p.BeforeTimestamp, 0)
	list, err := s.store.FindHistoryByID(ctx, p.ID, p.Limit, ts.Format("2006-01-02 15:04:05.999999"))
	if err != nil {
		return nil, fmt.Errorf("error finding last conversations from db: %w", err)
	}

	response := &chat.ChatHistoryResponse{
		ID:    p.ID,
		Count: len(list),
	}
	return response, nil
}

func (s *Service) SearchInConversation(ctx context.Context, p *chat.SearchInConversationPayload) (*chat.ChatHistoryResponse, error) {
	return nil, nil
}

func toConversationResources(conversations []models.Conversation) []*chat.Conversation {
	var conversationResources []*chat.Conversation
	for _, conv := range conversations {
		conversation := &chat.Conversation{
			ID:                     conv.ID,
			LastMessageSenderID:    conv.LastMessageSenderID,
			LastMessageContent:     conv.LastMessageContent,
			LastMessageDeliveredAt: conv.LastMessageDeliveredAt,
		}
		conversationResources = append(conversationResources, conversation)
	}
	return conversationResources
}
func (s *Service) GetConversationsList(ctx context.Context, p *chat.GetConversationsListPayload) (*chat.ConversationsListResponse, error) {
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting user from context: %w", err)
	}

	list, err := s.store.ListConversationsByUserID(ctx, u.ID, p.Limit)
	if err != nil {
		return nil, fmt.Errorf("error finding last conversations from db: %w", err)
	}

	resources := toConversationResources(list)
	response := &chat.ConversationsListResponse{
		Total:     len(list),
		Resources: resources,
	}
	return response, nil
}

func (s *Service) AddConversation(ctx context.Context, p *chat.AddConversationPayload) (*chat.OperationStatusResponse, error) {
	return nil, nil
}
