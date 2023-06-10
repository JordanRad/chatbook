package auth

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/JordanRad/chatbook/services/internal/gen/user"
)

// User struct, used across the application
type User struct {
	ID, Email, FirstName, LastName, Password string
	FriendsList                              FriendsList
}

// Workout Exercise Collection Type
type FriendsList []*user.Friend

// Value implements the driver Value method
func (fl FriendsList) Value() (driver.Value, error) {
	return json.Marshal(fl)
}

// Scan implements the Scan method
func (fl *FriendsList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(b, &fl)
	return err
}
