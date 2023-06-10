package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/encryption"
	"github.com/JordanRad/chatbook/services/internal/gen/user"
)

type Service struct {
	store   auth.UserStore
	encrypt encryption.Encryption
	logger  *log.Logger
}

type LoginDetails struct {
	Email    string
	Password string
}

func NewService(store auth.UserStore,
	encryption encryption.Encryption,
	logger *log.Logger) *Service {
	return &Service{
		store:   store,
		encrypt: encryption,
		logger:  logger,
	}
}

// Compile time assertion that this service implements the generated interface
var _ user.Service = (*Service)(nil)

func areNamesValid(fname, lname string) bool {
	return fname != "" && lname != ""
}

func (s *Service) Register(ctx context.Context, p *user.RegisterPayload) (*user.RegisterResponse, error) {
	if p.Password != p.ConfirmedPassword {
		return nil, user.MakeUnmatchingPassowrds(errors.New("Passwords do not match"))
	}

	encryptedPassword, err := s.encrypt.EncryptPassword(p.Password)
	if err != nil {
		return nil, fmt.Errorf("error encrypting password: %w", err)
	}

	newUser := &auth.User{
		Email:     p.Email,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Password:  encryptedPassword,
	}

	_, err = s.store.Register(ctx, newUser)
	if err != nil {
		return nil, fmt.Errorf("error creating new user: %w", err)
	}

	response := &user.RegisterResponse{
		Message: "User has been registered successfully",
	}

	return response, nil
}

func (s *Service) GetProfile(ctx context.Context) (*user.UserProfileResponse, error) {
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting user from context: %w", err)
	}
	result, err := s.store.GetUserByEmail(ctx, u.Email)

	if err != nil {
		return nil, fmt.Errorf("error getting user profile: %w", err)
	}

	response := &user.UserProfileResponse{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Email:       result.Email,
		FriendsList: result.FriendsList,
	}
	return response, nil
}

func (s *Service) UpdateProfileNames(ctx context.Context, p *user.UpdateProfileNamesPayload) (*user.OperationStatusResponse, error) {
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	if !areNamesValid(p.FirstName, p.LastName) {
		return nil, fmt.Errorf("names payload is incomplete")
	}

	err = s.store.UpdateProfileNames(ctx, u.ID, p.FirstName, p.LastName)
	if err != nil {
		return nil, fmt.Errorf("error updating user settings: %w", err)
	}

	r := &user.OperationStatusResponse{
		Message: "Profile names have been updated successfully",
	}
	return r, nil
}

func (s *Service) AddFriend(ctx context.Context, p *user.AddFriendPayload) (res *user.OperationStatusResponse, err error) {
	return nil, nil
}

func (s *Service) RemoveFriend(ctx context.Context, p *user.RemoveFriendPayload) (res *user.OperationStatusResponse, err error) {
	return nil, nil
}
