package dbchat

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JordanRad/chatbook/services/cmd/chat-service/chat"
	"github.com/JordanRad/chatbook/services/cmd/chat-service/db/models"
	"github.com/JordanRad/chatbook/services/internal/auth"
)

type Store struct {
	DB *sql.DB
}

var _ chat.Store = (*Store)(nil)

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

func (s *Store) FindHistoryByID(ctx context.Context, ID string, limit int, beforeTS string) ([]models.ConversationMessage, error) {
	rows, err := s.DB.Query(`
		SELECT sender_id, content, ts FROM messages
		WHERE conversation_id = '$1
		AND ts < $2
		ORDER BY ts DESC LIMIT $3;
		`, ID, beforeTS, limit)
	if err != nil {
		return nil, fmt.Errorf("error extracting chat history: %w", err)
	}
	defer rows.Close()

	messageList := make([]models.ConversationMessage, 0)
	for rows.Next() {
		message := &models.ConversationMessage{}

		err := rows.Scan(
			&message.SenderID,
			&message.Content,
			&message.TS)

		if err != nil {
			return nil, fmt.Errorf("error mapping a message row: %w", err)
		}

		messageList = append(messageList, *message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return messageList, nil
}

func (s *Store) FindBySearchInput(ctx context.Context, ID, input string, limit int) ([]models.ConversationMessage, error) {
	rows, err := s.DB.Query(`
		SELECT sender_id, content, ts
		FROM messages
		WHERE to_tsvector('english', content) @@ to_tsquery('english', $1)
		AND conversation_id = $2
		ORDER BY ts DESC LIMIT $3;
		`, input, ID, limit)
	if err != nil {
		return nil, fmt.Errorf("error extracting messages using full-text-search: %w", err)
	}
	defer rows.Close()

	messageList := make([]models.ConversationMessage, 0)
	for rows.Next() {
		message := &models.ConversationMessage{}

		err := rows.Scan(
			&message.SenderID,
			&message.Content,
			&message.TS)

		if err != nil {
			return nil, fmt.Errorf("error mapping a message row: %w", err)
		}

		messageList = append(messageList, *message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return messageList, nil
}
func (s *Store) ListConversationsByUserID(ctx context.Context, userID string, limit int) ([]models.Conversation, error) {
	rows, err := s.DB.Query(`
		WITH conversations_by_id_and_ts AS (
			SELECT m.conversation_id as conversation_id ,MAX(ts) AS max_ts
			FROM messages m
			JOIN conversations_participants cp ON m.conversation_id = cp.conversation_id
			WHERE cp.participant_id = $1
			GROUP BY m.conversation_id
			LIMIT $2
		)
		
		SELECT m.conversation_id as id, m.sender_id as last_message_sender_id , m.content last_message_content, m.ts as last_message_delivered_at
		FROM messages m
		JOIN conversations_by_id_and_ts as subset ON m.conversation_id = subset.conversation_id
		WHERE m.ts = subset.max_ts
		ORDER BY subset.max_ts DESC;
		`, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("error extracting last conversations list: %w", err)
	}
	defer rows.Close()

	messageList := make([]models.Conversation, 0)
	for rows.Next() {
		message := &models.Conversation{}

		err := rows.Scan(
			&message.ID,
			&message.LastMessageSenderID,
			&message.LastMessageContent,
			&message.LastMessageDeliveredAt)

		if err != nil {
			return nil, fmt.Errorf("error mapping a conversation row: %w", err)
		}

		messageList = append(messageList, *message)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return messageList, nil
}
func (s *Store) CreateConversation(ctx context.Context, participants []auth.FriendsList) error {
	return nil
}
