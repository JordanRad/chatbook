package dbchat

import (
	"context"
	"database/sql"

	"github.com/JordanRad/chatbook/services/cmd/chat-service/chat"
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

func (s *Store) GetHistory(ctx context.Context) error {
	return nil
}
