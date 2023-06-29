package auth

import (
	"context"
	"database/sql"
	"fmt"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . UserStore
type UserStore interface {
	GetAuthUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	Register(ctx context.Context, user *User) (*User, error)
	UpdateProfileNames(ctx context.Context, userID, firstName, lastName string) (User, error)
	AddFriend(ctx context.Context, userID, friendID string) error
	RemoveFriend(ctx context.Context, userID, friendID string) error
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
func (s *Store) GetAuthUserByEmail(ctx context.Context, email string) (*User, error) {
	result := new(User)

	row := s.DB.QueryRow(`
		SELECT u.id, u.email, u.first_name, u.last_name 
		FROM users u 
		WHERE u.email = $1
	`, email)
	err := row.Scan(
		&result.ID,
		&result.Email,
		&result.FirstName,
		&result.LastName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ErrUserNotFound{}
		}
		return nil, fmt.Errorf("error retrieving an auth user: %w", err)
	}

	return result, nil
}

// GetUserByEmail method returns a user from the database, based on a given email.
func (s *Store) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	result := new(User)

	row := s.DB.QueryRow(`
		WITH searched_user AS (
			SELECT id
			FROM users
			WHERE email = $1
		)
		SELECT
			u.id,
			u.first_name,
			u.last_name,
			u.email,
			u.password,
			COALESCE((
				SELECT JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', uf.id,
						'firstName', uf.first_name,
						'lastName', uf.last_name,
						'email', uf.email
					)
				)
				FROM users uf
				WHERE uf.id IN (
					SELECT CASE
						WHEN f.user_id = searched_user.id THEN f.friend_id
						ELSE f.user_id
					END
					FROM friendships f
					WHERE f.user_id = searched_user.id OR f.friend_id = searched_user.id
				) AND uf.id != searched_user.id
			), '[]'::json) AS friendsList
		FROM users u
		CROSS JOIN searched_user
		WHERE u.id = searched_user.id;
	`, email)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Password,
		&result.FriendsList,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ErrUserNotFound{}
		}
		return nil, fmt.Errorf("error retrieving a workout: %w", err)
	}

	return result, nil
}

// Register method add a new user to the database and returns the insterted user back along with the errors, if any error is present.
func (s *Store) Register(ctx context.Context, user *User) (*User, error) {
	var uid string
	err := s.DB.QueryRow(`INSERT INTO users(email, first_name, last_name, password) VALUES ($1,$2,$3,$4) RETURNING id;`, user.Email, user.FirstName, user.LastName, user.Password).Scan(&uid)
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
func (s *Store) UpdateProfileNames(ctx context.Context, userID, fname, lname string) (User, error) {
	_, err := s.DB.Exec(`UPDATE users SET first_name = $1,last_name = $2 WHERE id = $3;`, fname, lname, userID)
	if err != nil {
		return User{}, fmt.Errorf("error updating profile names: %w", err)
	}

	updatedUser := User{}
	err = s.DB.QueryRow(`SELECT id, first_name, last_name, email FROM users WHERE id = $1;`, userID).Scan(&updatedUser.ID, &updatedUser.FirstName, &updatedUser.LastName, &updatedUser.Email)
	if err != nil {
		return User{}, fmt.Errorf("error fetching updated user: %w", err)
	}

	return updatedUser, nil
}

func (s *Store) AddFriend(ctx context.Context, userID, friendID string) error {
	_, err := s.DB.Exec(`INSERT INTO friendships (user_id, friend_id) VALUES ($1 , $2)`, userID, friendID)
	if err != nil {
		return fmt.Errorf("error inserting a new friend in db: %w", err)
	}
	return nil

}
func (s *Store) RemoveFriend(ctx context.Context, userID, friendID string) error {
	_, err := s.DB.Exec(`DELETE from friendships WHERE (user_id = $1 AND friend_id = $2) OR (user_id = $2 AND friend_id = $1)`, userID, friendID)
	if err != nil {
		return fmt.Errorf("error deleting a friend from db: %w", err)
	}
	return nil
}
