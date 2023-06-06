package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Custom Token claims, used for the whole
// Smart Fit application
type TokenClaims struct {
	UserID string
	Email  string
	Exp    float64
}

// JWT secret used to create the signature
var secret = []byte(os.Getenv("CHATBOOK_JWT_SECRET"))

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate
//counterfeiter:generate . JWTClient
type JWTClient interface {
	GenerateJWT(userID, email string, refresh bool) (string, error)
	ExtractJWTCLaims(tokenString string) (TokenClaims, error)
	ValidateJWT(tokenString string) (bool, error)
}

type JWTService struct{}

// Compile time assertion that this service implements the generated interface
var _ JWTClient = (*JWTService)(nil)

// This method generates a valid JSON Web Token (JWT)
func (s *JWTService) GenerateJWT(userID, email string, refresh bool) (string, error) {
	exp := time.Hour * 24

	if refresh {
		exp = time.Hour * 72
	}
	// Set custom claims to JWT
	claims := &jwt.MapClaims{
		"iss": "chatbook-services-vhvfhsv1e-580b-4404-a187-0f7lfv24dfs-vfba-c64vbmtubdku-5a-9e90a02bb",
		"exp": time.Now().Add(exp).Unix(),
		"data": map[string]string{
			"userID": userID,
			"email":  email,
		},
	}

	// Sign the token and add the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("error generating JWT: %w", err)
	}

	return tokenString, nil
}

// This methods returns custom token claims in a Go struct,
// based on a given JWT
func (s *JWTService) ExtractJWTCLaims(tokenString string) (TokenClaims, error) {
	//Parse the JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return TokenClaims{}, fmt.Errorf("error casting claims: %w", err)
	}

	// Extract the token's claims
	claims := token.Claims.(jwt.MapClaims)

	// Parse the claims to a struct
	exp := claims["exp"].(float64)
	data := claims["data"].(map[string]interface{})
	userID := data["userID"].(string)
	email := data["email"].(string)

	tokenClaims := TokenClaims{
		UserID: userID,
		Email:  email,
		Exp:    exp,
	}

	return tokenClaims, nil
}

// This method ensures that a given JWT is valid
func (s *JWTService) ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return secret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, fmt.Errorf("error signing the token: %w", err)
		}
	}
	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}
	return true, nil
}
