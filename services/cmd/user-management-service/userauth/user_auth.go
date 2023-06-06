package userauth

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/encryption"
	"github.com/JordanRad/chatbook/services/internal/auth/jwt"
	authgen "github.com/JordanRad/chatbook/services/internal/gen/auth"
)

type Service struct {
	store   auth.UserStore
	encrypt encryption.Encryption
	jwt     jwt.JWTClient
	logger  *log.Logger
}

type LoginDetails struct {
	Email    string
	Password string
}

func NewService(store auth.UserStore,
	encrypt encryption.Encryption,
	jwt jwt.JWTClient,
	logger *log.Logger) *Service {
	return &Service{
		store:   store,
		encrypt: encrypt,
		jwt:     jwt,
		logger:  logger,
	}
}

// Compile time assertion that this service implements the generated interface
var _ = (*Service)(nil)

func (s *Service) Login(ctx context.Context, p *authgen.LoginPayload) (*authgen.LoginResponse, error) {
	foundUser, err := s.store.GetUserByEmail(ctx, p.Email)

	if err != nil {
		err := authgen.MakeWrongCredentials(err)
		return nil, err
	}

	isPasswordCorrect := s.encrypt.CheckPassword(foundUser.Password, p.Password)
	if !isPasswordCorrect {
		err := authgen.MakeWrongCredentials(errors.New("password is not correct"))
		return nil, err
	}

	t, err := s.jwt.GenerateJWT(foundUser.ID, foundUser.Email, false)
	if err != nil {
		return nil, fmt.Errorf("error extracting jwt: %w", err)
	}

	rt, err := s.jwt.GenerateJWT(foundUser.ID, foundUser.Email, true)
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	response := &authgen.LoginResponse{
		Email:        p.Email,
		Token:        t,
		RefreshToken: rt,
	}
	return response, nil
}

func (s *Service) RefreshToken(ctx context.Context, p *authgen.RefreshTokenPayload) (*authgen.LoginResponse, error) {
	u, err := s.store.GetUserByEmail(ctx, p.Email)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user from the database: %w", err)
	}

	_, err = s.jwt.ValidateJWT(p.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("error validating refresh token: %w", err)
	}

	t, err := s.jwt.GenerateJWT(u.ID, u.Email, false)
	if err != nil {
		return nil, fmt.Errorf("error generating token: %w", err)
	}

	rt, err := s.jwt.GenerateJWT(u.ID, u.Email, true)
	if err != nil {
		return nil, fmt.Errorf("error generating refresh token: %w", err)
	}

	r := &authgen.LoginResponse{
		Email:        u.Email,
		Token:        t,
		RefreshToken: rt,
		ID:           nil,
	}
	return r, nil
}
