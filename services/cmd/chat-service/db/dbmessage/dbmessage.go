package dbmessage

import (
	"database/sql"
	"fmt"

	websocketsserver "github.com/JordanRad/chatbook/services/internal/websockets/websockets_server"
)

type Store struct {
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

var _ websocketsserver.Store = (*Store)(nil)

func (s *Store) SaveConversationMessage(ID, senderID, content string) error {
	result, err := s.DB.Exec(`INSERT INTO messages(conversation_id, sender_id, content) VALUES ($1,$2,$3);`, ID, senderID, content)
	if err != nil {
		return fmt.Errorf("error inserting new message entry: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		return nil
	}

	return fmt.Errorf("error inserting new message entry")
}
