package auth

import (
	"context"
	"database/sql"
	"fmt"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . UserStore
type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Register(ctx context.Context, user *User) (*User, error)
	UpdateProfileNames(ctx context.Context, userID, firstName, lastName string) error
}

type Store struct {
	DB *sql.DB
}

var _ UserStore = (*Store)(nil)

func NewStore(db *sql.DB) *Store {
	return &Store{
		DB: db,
	}
}

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "Such user has not been found"
}

// GetUserByEmail method returns a user from the database, based on a given email.
func (s *Store) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	result := new(User)

	row := s.DB.QueryRow(`SELECT id, firstname, lastname, email, password FROM users WHERE email = $1;`, email)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Password,
	)
	if err == sql.ErrNoRows {
		return nil, &ErrUserNotFound{}
	}

	return result, nil
}

// Register method add a new user to the database and returns the insterted user back along with the errors, if any error is present.
func (s *Store) Register(ctx context.Context, user *User) (*User, error) {
	var uid string
	err := s.DB.QueryRow(`INSERT INTO users(email, firstName, lastName, password) VALUES ($1,$2,$3,$4) RETURNING id;`, user.Email, user.FirstName, user.LastName, user.Password).Scan(&uid)
	if err != nil {
		return nil, fmt.Errorf("error registering new user: %w", err)
	}
	u := &User{
		ID:        uid,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return u, nil
}

// UpdateProfileNames method updates user's names in the database and returns error, if any error is present.
func (s *Store) UpdateProfileNames(ctx context.Context, userID, fname, lname string) error {
	result, err := s.DB.Exec(`UPDATE users SET firstName=$1,lastName=$2 WHERE id = $3;`, fname, lname, userID)

	if err != nil {
		return fmt.Errorf("error updating profile names: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		return nil
	}

	return fmt.Errorf("error updating profile names, zero rows affected")
}
