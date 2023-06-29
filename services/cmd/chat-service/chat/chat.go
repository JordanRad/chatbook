package chat

import (
	"context"
	"fmt"
	"log"

	"github.com/JordanRad/chatbook/services/cmd/chat-service/db/models"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/gen/chat"
)

type Service struct {
	logger *log.Logger
	store  Store
}

type Store interface {
	FindHistoryByID(ctx context.Context, ID, beforeTS string, limit int) ([]models.ConversationMessage, error)
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

func toMessagesList(conversationMessages []models.ConversationMessage) []*chat.ConversationMessage {
	var resources []*chat.ConversationMessage
	for _, m := range conversationMessages {
		r := &chat.ConversationMessage{
			SenderID:  m.SenderID,
			Timestamp: m.TS.String(),
			Content:   m.Content,
		}
		resources = append(resources, r)
	}
	return resources
}

func (s *Service) GetConversationHistory(ctx context.Context, p *chat.GetConversationHistoryPayload) (*chat.ChatHistoryResponse, error) {
	_, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting user from context: %w", err)
	}

	list, err := s.store.FindHistoryByID(ctx, p.ID, p.BeforeTimestamp, p.Limit)
	if err != nil {
		return nil, fmt.Errorf("error finding last conversations from db: %w", err)
	}

	resources := toMessagesList(list)

	response := &chat.ChatHistoryResponse{
		ID:       p.ID,
		Count:    len(list),
		Messages: resources,
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
			OtherParticipantID:     conv.OtherParticipantID,
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
