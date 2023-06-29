// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chat HTTP client types
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package client

import (
	chat "github.com/JordanRad/chatbook/services/internal/gen/chat"
	goa "goa.design/goa/v3/pkg"
)

// AddConversationRequestBody is the type of the "chat" service
// "addConversation" endpoint HTTP request body.
type AddConversationRequestBody struct {
	// Participants
	Participants []*FriendRequestBody `form:"participants" json:"participants" xml:"participants"`
}

// GetConversationHistoryResponseBody is the type of the "chat" service
// "getConversationHistory" endpoint HTTP response body.
type GetConversationHistoryResponseBody struct {
	// Chatroom ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Messages Count
	Count *int `form:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// Chat history
	Messages []*ConversationMessageResponseBody `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
}

// SearchInConversationResponseBody is the type of the "chat" service
// "searchInConversation" endpoint HTTP response body.
type SearchInConversationResponseBody struct {
	// Chatroom ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Messages Count
	Count *int `form:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// Chat history
	Messages []*ConversationMessageResponseBody `form:"messages,omitempty" json:"messages,omitempty" xml:"messages,omitempty"`
}

// GetConversationsListResponseBody is the type of the "chat" service
// "getConversationsList" endpoint HTTP response body.
type GetConversationsListResponseBody struct {
	// Messages Count
	Total *int `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// Chat history
	Resources []*ConversationResponseBody `form:"resources,omitempty" json:"resources,omitempty" xml:"resources,omitempty"`
}

// AddConversationResponseBody is the type of the "chat" service
// "addConversation" endpoint HTTP response body.
type AddConversationResponseBody struct {
	// Operation status
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// ConversationMessageResponseBody is used to define fields on response body
// types.
type ConversationMessageResponseBody struct {
	// Sender ID
	SenderID *string `form:"senderID,omitempty" json:"senderID,omitempty" xml:"senderID,omitempty"`
	// Timestamp of the message
	Timestamp *string `form:"timestamp,omitempty" json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// Message Content
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
}

// ConversationResponseBody is used to define fields on response body types.
type ConversationResponseBody struct {
	// Conversation ID
	ID *string `form:"ID,omitempty" json:"ID,omitempty" xml:"ID,omitempty"`
	// Sender ID
	LastMessageSenderID *string `form:"lastMessageSenderID,omitempty" json:"lastMessageSenderID,omitempty" xml:"lastMessageSenderID,omitempty"`
	// Last message
	LastMessageContent *string `form:"lastMessageContent,omitempty" json:"lastMessageContent,omitempty" xml:"lastMessageContent,omitempty"`
	// TS for delivered time
	LastMessageDeliveredAt *string `form:"lastMessageDeliveredAt,omitempty" json:"lastMessageDeliveredAt,omitempty" xml:"lastMessageDeliveredAt,omitempty"`
	// TS for delivered time
	OtherParticipantID *string `form:"otherParticipantID,omitempty" json:"otherParticipantID,omitempty" xml:"otherParticipantID,omitempty"`
}

// FriendRequestBody is used to define fields on request body types.
type FriendRequestBody struct {
	// User ID
	ID string `form:"id" json:"id" xml:"id"`
	// Email
	Email string `form:"email" json:"email" xml:"email"`
	// First name
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
}

// NewAddConversationRequestBody builds the HTTP request body from the payload
// of the "addConversation" endpoint of the "chat" service.
func NewAddConversationRequestBody(p *chat.AddConversationPayload) *AddConversationRequestBody {
	body := &AddConversationRequestBody{}
	if p.Participants != nil {
		body.Participants = make([]*FriendRequestBody, len(p.Participants))
		for i, val := range p.Participants {
			body.Participants[i] = marshalChatFriendToFriendRequestBody(val)
		}
	}
	return body
}

// NewGetConversationHistoryChatHistoryResponseOK builds a "chat" service
// "getConversationHistory" endpoint result from a HTTP "OK" response.
func NewGetConversationHistoryChatHistoryResponseOK(body *GetConversationHistoryResponseBody) *chat.ChatHistoryResponse {
	v := &chat.ChatHistoryResponse{
		ID:    *body.ID,
		Count: *body.Count,
	}
	v.Messages = make([]*chat.ConversationMessage, len(body.Messages))
	for i, val := range body.Messages {
		v.Messages[i] = unmarshalConversationMessageResponseBodyToChatConversationMessage(val)
	}

	return v
}

// NewSearchInConversationChatHistoryResponseOK builds a "chat" service
// "searchInConversation" endpoint result from a HTTP "OK" response.
func NewSearchInConversationChatHistoryResponseOK(body *SearchInConversationResponseBody) *chat.ChatHistoryResponse {
	v := &chat.ChatHistoryResponse{
		ID:    *body.ID,
		Count: *body.Count,
	}
	v.Messages = make([]*chat.ConversationMessage, len(body.Messages))
	for i, val := range body.Messages {
		v.Messages[i] = unmarshalConversationMessageResponseBodyToChatConversationMessage(val)
	}

	return v
}

// NewGetConversationsListConversationsListResponseOK builds a "chat" service
// "getConversationsList" endpoint result from a HTTP "OK" response.
func NewGetConversationsListConversationsListResponseOK(body *GetConversationsListResponseBody) *chat.ConversationsListResponse {
	v := &chat.ConversationsListResponse{
		Total: *body.Total,
	}
	v.Resources = make([]*chat.Conversation, len(body.Resources))
	for i, val := range body.Resources {
		v.Resources[i] = unmarshalConversationResponseBodyToChatConversation(val)
	}

	return v
}

// NewAddConversationOperationStatusResponseOK builds a "chat" service
// "addConversation" endpoint result from a HTTP "OK" response.
func NewAddConversationOperationStatusResponseOK(body *AddConversationResponseBody) *chat.OperationStatusResponse {
	v := &chat.OperationStatusResponse{
		Message: *body.Message,
	}

	return v
}

// ValidateGetConversationHistoryResponseBody runs the validations defined on
// GetConversationHistoryResponseBody
func ValidateGetConversationHistoryResponseBody(body *GetConversationHistoryResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("count", "body"))
	}
	if body.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("messages", "body"))
	}
	for _, e := range body.Messages {
		if e != nil {
			if err2 := ValidateConversationMessageResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSearchInConversationResponseBody runs the validations defined on
// SearchInConversationResponseBody
func ValidateSearchInConversationResponseBody(body *SearchInConversationResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Count == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("count", "body"))
	}
	if body.Messages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("messages", "body"))
	}
	for _, e := range body.Messages {
		if e != nil {
			if err2 := ValidateConversationMessageResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateGetConversationsListResponseBody runs the validations defined on
// GetConversationsListResponseBody
func ValidateGetConversationsListResponseBody(body *GetConversationsListResponseBody) (err error) {
	if body.Total == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("total", "body"))
	}
	if body.Resources == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("resources", "body"))
	}
	for _, e := range body.Resources {
		if e != nil {
			if err2 := ValidateConversationResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateAddConversationResponseBody runs the validations defined on
// AddConversationResponseBody
func ValidateAddConversationResponseBody(body *AddConversationResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateConversationMessageResponseBody runs the validations defined on
// ConversationMessageResponseBody
func ValidateConversationMessageResponseBody(body *ConversationMessageResponseBody) (err error) {
	if body.Timestamp == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timestamp", "body"))
	}
	if body.SenderID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("senderID", "body"))
	}
	if body.Content == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("content", "body"))
	}
	return
}

// ValidateConversationResponseBody runs the validations defined on
// ConversationResponseBody
func ValidateConversationResponseBody(body *ConversationResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ID", "body"))
	}
	if body.LastMessageSenderID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastMessageSenderID", "body"))
	}
	if body.LastMessageContent == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastMessageContent", "body"))
	}
	if body.LastMessageDeliveredAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("lastMessageDeliveredAt", "body"))
	}
	if body.OtherParticipantID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("otherParticipantID", "body"))
	}
	return
}
