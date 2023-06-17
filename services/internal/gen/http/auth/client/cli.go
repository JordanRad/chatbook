// Code generated by goa v3.11.3, DO NOT EDIT.
//
// auth HTTP client CLI support package
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package client

import (
	"encoding/json"
	"fmt"

	auth "github.com/JordanRad/chatbook/services/internal/gen/auth"
)

// BuildRefreshTokenPayload builds the payload for the auth refreshToken
// endpoint from CLI flags.
func BuildRefreshTokenPayload(authRefreshTokenBody string) (*auth.RefreshTokenPayload, error) {
	var err error
	var body RefreshTokenRequestBody
	{
		err = json.Unmarshal([]byte(authRefreshTokenBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Natus quisquam.\",\n      \"refreshToken\": \"Illo omnis nesciunt minus sed.\"\n   }'")
		}
	}
	v := &auth.RefreshTokenPayload{
		Email:        body.Email,
		RefreshToken: body.RefreshToken,
	}

	return v, nil
}

// BuildLoginPayload builds the payload for the auth login endpoint from CLI
// flags.
func BuildLoginPayload(authLoginBody string) (*auth.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(authLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Labore tempora.\",\n      \"password\": \"Ut ut.\"\n   }'")
		}
	}
	v := &auth.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}
