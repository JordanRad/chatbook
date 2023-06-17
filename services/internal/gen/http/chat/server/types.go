// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chat HTTP server types
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package server

import (
	chat "github.com/JordanRad/chatbook/services/internal/gen/chat"
	goa "goa.design/goa/v3/pkg"
)

// GetConversationsListRequestBody is the type of the "chat" service
// "getConversationsList" endpoint HTTP request body.
type GetConversationsListRequestBody struct {
	// Conversation ID
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
}

// AddConversationRequestBody is the type of the "chat" service
// "addConversation" endpoint HTTP request body.
type AddConversationRequestBody struct {
	// Participants
	Participants []*FriendRequestBody `form:"participants,omitempty" json:"participants,omitempty" xml:"participants,omitempty"`
}

// GetConversationHistoryResponseBody is the type of the "chat" service
// "getConversationHistory" endpoint HTTP response body.
type GetConversationHistoryResponseBody struct {
	// Chatroom ID
	ID string `form:"id" json:"id" xml:"id"`
	// Messages Count
	Count int `form:"count" json:"count" xml:"count"`
	// Chat history
	Messages []*ConversationMessageResponseBody `form:"messages" json:"messages" xml:"messages"`
}

// SearchInConversationResponseBody is the type of the "chat" service
// "searchInConversation" endpoint HTTP response body.
type SearchInConversationResponseBody struct {
	// Chatroom ID
	ID string `form:"id" json:"id" xml:"id"`
	// Messages Count
	Count int `form:"count" json:"count" xml:"count"`
	// Chat history
	Messages []*ConversationMessageResponseBody `form:"messages" json:"messages" xml:"messages"`
}

// GetConversationsListResponseBody is the type of the "chat" service
// "getConversationsList" endpoint HTTP response body.
type GetConversationsListResponseBody struct {
	// Messages Count
	Total int `form:"total" json:"total" xml:"total"`
	// Chat history
	Resources []*ConversationResponseBody `form:"resources" json:"resources" xml:"resources"`
}

// AddConversationResponseBody is the type of the "chat" service
// "addConversation" endpoint HTTP response body.
type AddConversationResponseBody struct {
	// Operation status
	Message string `form:"message" json:"message" xml:"message"`
}

// ConversationMessageResponseBody is used to define fields on response body
// types.
type ConversationMessageResponseBody struct {
	// Sender ID
	SenderID string `form:"senderID" json:"senderID" xml:"senderID"`
	// Timestamp of the message
	Timestamp float64 `form:"timestamp" json:"timestamp" xml:"timestamp"`
	// Message Content
	Content string `form:"content" json:"content" xml:"content"`
}

// ConversationResponseBody is used to define fields on response body types.
type ConversationResponseBody struct {
	// Conversation ID
	ID string `form:"ID" json:"ID" xml:"ID"`
	// Timestamp of the message
	Participants float64 `form:"participants" json:"participants" xml:"participants"`
	// Last message
	LastMessage string `form:"lastMessage" json:"lastMessage" xml:"lastMessage"`
	// TS for delivered time
	DeliveredAt int64 `form:"deliveredAt" json:"deliveredAt" xml:"deliveredAt"`
}

// FriendRequestBody is used to define fields on request body types.
type FriendRequestBody struct {
	// User ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First name
	FirstName *string `form:"firstName,omitempty" json:"firstName,omitempty" xml:"firstName,omitempty"`
	// Last name
	LastName *string `form:"lastName,omitempty" json:"lastName,omitempty" xml:"lastName,omitempty"`
}

// NewGetConversationHistoryResponseBody builds the HTTP response body from the
// result of the "getConversationHistory" endpoint of the "chat" service.
func NewGetConversationHistoryResponseBody(res *chat.ChatHistoryResponse) *GetConversationHistoryResponseBody {
	body := &GetConversationHistoryResponseBody{
		ID:    res.ID,
		Count: res.Count,
	}
	if res.Messages != nil {
		body.Messages = make([]*ConversationMessageResponseBody, len(res.Messages))
		for i, val := range res.Messages {
			body.Messages[i] = marshalChatConversationMessageToConversationMessageResponseBody(val)
		}
	}
	return body
}

// NewSearchInConversationResponseBody builds the HTTP response body from the
// result of the "searchInConversation" endpoint of the "chat" service.
func NewSearchInConversationResponseBody(res *chat.ChatHistoryResponse) *SearchInConversationResponseBody {
	body := &SearchInConversationResponseBody{
		ID:    res.ID,
		Count: res.Count,
	}
	if res.Messages != nil {
		body.Messages = make([]*ConversationMessageResponseBody, len(res.Messages))
		for i, val := range res.Messages {
			body.Messages[i] = marshalChatConversationMessageToConversationMessageResponseBody(val)
		}
	}
	return body
}

// NewGetConversationsListResponseBody builds the HTTP response body from the
// result of the "getConversationsList" endpoint of the "chat" service.
func NewGetConversationsListResponseBody(res *chat.ConversationsListResponse) *GetConversationsListResponseBody {
	body := &GetConversationsListResponseBody{
		Total: res.Total,
	}
	if res.Resources != nil {
		body.Resources = make([]*ConversationResponseBody, len(res.Resources))
		for i, val := range res.Resources {
			body.Resources[i] = marshalChatConversationToConversationResponseBody(val)
		}
	}
	return body
}

// NewAddConversationResponseBody builds the HTTP response body from the result
// of the "addConversation" endpoint of the "chat" service.
func NewAddConversationResponseBody(res *chat.OperationStatusResponse) *AddConversationResponseBody {
	body := &AddConversationResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetConversationHistoryPayload builds a chat service
// getConversationHistory endpoint payload.
func NewGetConversationHistoryPayload(id string, limit string, beforeTimestamp int64) *chat.GetConversationHistoryPayload {
	v := &chat.GetConversationHistoryPayload{}
	v.ID = id
	v.Limit = limit
	v.BeforeTimestamp = beforeTimestamp

	return v
}

// NewSearchInConversationPayload builds a chat service searchInConversation
// endpoint payload.
func NewSearchInConversationPayload(id string, limit string, searchInput string) *chat.SearchInConversationPayload {
	v := &chat.SearchInConversationPayload{}
	v.ID = id
	v.Limit = limit
	v.SearchInput = searchInput

	return v
}

// NewGetConversationsListPayload builds a chat service getConversationsList
// endpoint payload.
func NewGetConversationsListPayload(body *GetConversationsListRequestBody, limit string) *chat.GetConversationsListPayload {
	v := &chat.GetConversationsListPayload{
		ID: *body.ID,
	}
	v.Limit = limit

	return v
}

// NewAddConversationPayload builds a chat service addConversation endpoint
// payload.
func NewAddConversationPayload(body *AddConversationRequestBody) *chat.AddConversationPayload {
	v := &chat.AddConversationPayload{}
	v.Participants = make([]*chat.Friend, len(body.Participants))
	for i, val := range body.Participants {
		v.Participants[i] = unmarshalFriendRequestBodyToChatFriend(val)
	}

	return v
}

// ValidateGetConversationsListRequestBody runs the validations defined on
// GetConversationsListRequestBody
func ValidateGetConversationsListRequestBody(body *GetConversationsListRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ID", "body"))
	}
	return
}

// ValidateAddConversationRequestBody runs the validations defined on
// AddConversationRequestBody
func ValidateAddConversationRequestBody(body *AddConversationRequestBody) (err error) {
	if body.Participants == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("participants", "body"))
	}
	for _, e := range body.Participants {
		if e != nil {
			if err2 := ValidateFriendRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateFriendRequestBody runs the validations defined on FriendRequestBody
func ValidateFriendRequestBody(body *FriendRequestBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("firstName", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastName", "body"))
	}
	return
}
