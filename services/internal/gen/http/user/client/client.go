// Code generated by goa v3.11.3, DO NOT EDIT.
//
// user client HTTP transport
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the user service endpoint HTTP clients.
type Client struct {
	// Register Doer is the HTTP client used to make requests to the register
	// endpoint.
	RegisterDoer goahttp.Doer

	// GetProfile Doer is the HTTP client used to make requests to the getProfile
	// endpoint.
	GetProfileDoer goahttp.Doer

	// UpdateProfileNames Doer is the HTTP client used to make requests to the
	// updateProfileNames endpoint.
	UpdateProfileNamesDoer goahttp.Doer

	// AddFriend Doer is the HTTP client used to make requests to the addFriend
	// endpoint.
	AddFriendDoer goahttp.Doer

	// RemoveFriend Doer is the HTTP client used to make requests to the
	// removeFriend endpoint.
	RemoveFriendDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the user service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RegisterDoer:           doer,
		GetProfileDoer:         doer,
		UpdateProfileNamesDoer: doer,
		AddFriendDoer:          doer,
		RemoveFriendDoer:       doer,
		CORSDoer:               doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// Register returns an endpoint that makes HTTP requests to the user service
// register server.
func (c *Client) Register() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterRequest(c.encoder)
		decodeResponse = DecodeRegisterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRegisterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "register", err)
		}
		return decodeResponse(resp)
	}
}

// GetProfile returns an endpoint that makes HTTP requests to the user service
// getProfile server.
func (c *Client) GetProfile() goa.Endpoint {
	var (
		decodeResponse = DecodeGetProfileResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetProfileRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetProfileDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "getProfile", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateProfileNames returns an endpoint that makes HTTP requests to the user
// service updateProfileNames server.
func (c *Client) UpdateProfileNames() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateProfileNamesRequest(c.encoder)
		decodeResponse = DecodeUpdateProfileNamesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUpdateProfileNamesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateProfileNamesDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "updateProfileNames", err)
		}
		return decodeResponse(resp)
	}
}

// AddFriend returns an endpoint that makes HTTP requests to the user service
// addFriend server.
func (c *Client) AddFriend() goa.Endpoint {
	var (
		decodeResponse = DecodeAddFriendResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddFriendRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddFriendDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "addFriend", err)
		}
		return decodeResponse(resp)
	}
}

// RemoveFriend returns an endpoint that makes HTTP requests to the user
// service removeFriend server.
func (c *Client) RemoveFriend() goa.Endpoint {
	var (
		decodeResponse = DecodeRemoveFriendResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildRemoveFriendRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveFriendDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("user", "removeFriend", err)
		}
		return decodeResponse(resp)
	}
}
