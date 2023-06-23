// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chat client HTTP transport
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

// Client lists the chat service endpoint HTTP clients.
type Client struct {
	// GetConversationHistory Doer is the HTTP client used to make requests to the
	// getConversationHistory endpoint.
	GetConversationHistoryDoer goahttp.Doer

	// SearchInConversation Doer is the HTTP client used to make requests to the
	// searchInConversation endpoint.
	SearchInConversationDoer goahttp.Doer

	// GetConversationsList Doer is the HTTP client used to make requests to the
	// getConversationsList endpoint.
	GetConversationsListDoer goahttp.Doer

	// AddConversation Doer is the HTTP client used to make requests to the
	// addConversation endpoint.
	AddConversationDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the chat service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetConversationHistoryDoer: doer,
		SearchInConversationDoer:   doer,
		GetConversationsListDoer:   doer,
		AddConversationDoer:        doer,
		CORSDoer:                   doer,
		RestoreResponseBody:        restoreBody,
		scheme:                     scheme,
		host:                       host,
		decoder:                    dec,
		encoder:                    enc,
	}
}

// GetConversationHistory returns an endpoint that makes HTTP requests to the
// chat service getConversationHistory server.
func (c *Client) GetConversationHistory() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetConversationHistoryRequest(c.encoder)
		decodeResponse = DecodeGetConversationHistoryResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetConversationHistoryRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetConversationHistoryDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chat", "getConversationHistory", err)
		}
		return decodeResponse(resp)
	}
}

// SearchInConversation returns an endpoint that makes HTTP requests to the
// chat service searchInConversation server.
func (c *Client) SearchInConversation() goa.Endpoint {
	var (
		encodeRequest  = EncodeSearchInConversationRequest(c.encoder)
		decodeResponse = DecodeSearchInConversationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSearchInConversationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SearchInConversationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chat", "searchInConversation", err)
		}
		return decodeResponse(resp)
	}
}

// GetConversationsList returns an endpoint that makes HTTP requests to the
// chat service getConversationsList server.
func (c *Client) GetConversationsList() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetConversationsListRequest(c.encoder)
		decodeResponse = DecodeGetConversationsListResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildGetConversationsListRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetConversationsListDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chat", "getConversationsList", err)
		}
		return decodeResponse(resp)
	}
}

// AddConversation returns an endpoint that makes HTTP requests to the chat
// service addConversation server.
func (c *Client) AddConversation() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddConversationRequest(c.encoder)
		decodeResponse = DecodeAddConversationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildAddConversationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddConversationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("chat", "addConversation", err)
		}
		return decodeResponse(resp)
	}
}