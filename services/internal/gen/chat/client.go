// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chat client
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package chat

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "chat" service client.
type Client struct {
	GetConversationHistoryEndpoint goa.Endpoint
	SearchInConversationEndpoint   goa.Endpoint
	GetConversationsListEndpoint   goa.Endpoint
	AddConversationEndpoint        goa.Endpoint
}

// NewClient initializes a "chat" service client given the endpoints.
func NewClient(getConversationHistory, searchInConversation, getConversationsList, addConversation goa.Endpoint) *Client {
	return &Client{
		GetConversationHistoryEndpoint: getConversationHistory,
		SearchInConversationEndpoint:   searchInConversation,
		GetConversationsListEndpoint:   getConversationsList,
		AddConversationEndpoint:        addConversation,
	}
}

// GetConversationHistory calls the "getConversationHistory" endpoint of the
// "chat" service.
func (c *Client) GetConversationHistory(ctx context.Context, p *GetConversationHistoryPayload) (res *ChatHistoryResponse, err error) {
	var ires any
	ires, err = c.GetConversationHistoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatHistoryResponse), nil
}

// SearchInConversation calls the "searchInConversation" endpoint of the "chat"
// service.
func (c *Client) SearchInConversation(ctx context.Context, p *SearchInConversationPayload) (res *ChatHistoryResponse, err error) {
	var ires any
	ires, err = c.SearchInConversationEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ChatHistoryResponse), nil
}

// GetConversationsList calls the "getConversationsList" endpoint of the "chat"
// service.
func (c *Client) GetConversationsList(ctx context.Context, p *GetConversationsListPayload) (res *ConversationsListResponse, err error) {
	var ires any
	ires, err = c.GetConversationsListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ConversationsListResponse), nil
}

// AddConversation calls the "addConversation" endpoint of the "chat" service.
func (c *Client) AddConversation(ctx context.Context, p *AddConversationPayload) (res *OperationStatusResponse, err error) {
	var ires any
	ires, err = c.AddConversationEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*OperationStatusResponse), nil
}