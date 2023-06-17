package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/encryption"
	"github.com/JordanRad/chatbook/services/internal/gen/user"
	"google.golang.org/grpc"

	notificationsprotobuf "github.com/JordanRad/chatbook/services/internal/gen/grpc/notification/pb"
)

type NotificationService interface {
	NotifyUserNamesUpdate(ctx context.Context, in *notificationsprotobuf.NotifyUserNamesUpdateRequest, opts ...grpc.CallOption) (*notificationsprotobuf.NotifyUserNamesUpdateResponse, error)
}

type Service struct {
	store                auth.UserStore
	encrypt              encryption.Encryption
	logger               *log.Logger
	notificationsService notificationsprotobuf.NotificationClient
}

type LoginDetails struct {
	Email    string
	Password string
}

func NewService(store auth.UserStore,
	encryption encryption.Encryption,
	logger *log.Logger, notificationsService NotificationService) *Service {
	return &Service{
		store:                store,
		encrypt:              encryption,
		logger:               logger,
		notificationsService: notificationsService,
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

	updatedUser, err := s.store.UpdateProfileNames(ctx, u.ID, p.FirstName, p.LastName)
	if err != nil {
		return nil, fmt.Errorf("error updating user settings: %w", err)
	}

	grpcPayload := &notificationsprotobuf.NotifyUserNamesUpdateRequest{
		Id:           updatedUser.ID,
		FirstName:    updatedUser.FirstName,
		OldFirstName: u.FirstName,
		OldLastName:  u.LastName,
		LastName:     updatedUser.LastName,
		Ts:           time.Now().Format("24-07-2009"),
	}
	_, err = s.notificationsService.NotifyUserNamesUpdate(ctx, grpcPayload)
	if err != nil {
		return nil, fmt.Errorf("grpc error notifying chat service: %w", err)
	}

	r := &user.OperationStatusResponse{
		Message: "Profile names have been updated successfully",
	}
	return r, nil
}

func (s *Service) AddFriend(ctx context.Context, p *user.AddFriendPayload) (*user.OperationStatusResponse, error) {
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	err = s.store.AddFriend(ctx, u.ID, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error adding new friend: %w", err)
	}

	r := &user.OperationStatusResponse{
		Message: "Friend has been added successfully",
	}
	return r, nil
}

func (s *Service) RemoveFriend(ctx context.Context, p *user.RemoveFriendPayload) (*user.OperationStatusResponse, error) {
	u, err := auth.UserInContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	err = s.store.RemoveFriend(ctx, u.ID, p.ID)
	if err != nil {
		return nil, fmt.Errorf("error removing a friend: %w", err)
	}

	r := &user.OperationStatusResponse{
		Message: "Friend has been removed successfully",
	}
	return r, nil
}
