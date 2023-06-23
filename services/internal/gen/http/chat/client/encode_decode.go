// Code generated by goa v3.11.3, DO NOT EDIT.
//
// chat HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	chat "github.com/JordanRad/chatbook/services/internal/gen/chat"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetConversationHistoryRequest instantiates a HTTP request object with
// method and path set to call the "chat" service "getConversationHistory"
// endpoint
func (c *Client) BuildGetConversationHistoryRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*chat.GetConversationHistoryPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chat", "getConversationHistory", "*chat.GetConversationHistoryPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetConversationHistoryChatPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat", "getConversationHistory", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetConversationHistoryRequest returns an encoder for requests sent to
// the chat getConversationHistory server.
func EncodeGetConversationHistoryRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chat.GetConversationHistoryPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat", "getConversationHistory", "*chat.GetConversationHistoryPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("beforeTimestamp", fmt.Sprintf("%v", p.BeforeTimestamp))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetConversationHistoryResponse returns a decoder for responses
// returned by the chat getConversationHistory endpoint. restoreBody controls
// whether the response body should be restored after having been read.
func DecodeGetConversationHistoryResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetConversationHistoryResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat", "getConversationHistory", err)
			}
			err = ValidateGetConversationHistoryResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat", "getConversationHistory", err)
			}
			res := NewGetConversationHistoryChatHistoryResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat", "getConversationHistory", resp.StatusCode, string(body))
		}
	}
}

// BuildSearchInConversationRequest instantiates a HTTP request object with
// method and path set to call the "chat" service "searchInConversation"
// endpoint
func (c *Client) BuildSearchInConversationRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id string
	)
	{
		p, ok := v.(*chat.SearchInConversationPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("chat", "searchInConversation", "*chat.SearchInConversationPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SearchInConversationChatPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat", "searchInConversation", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSearchInConversationRequest returns an encoder for requests sent to
// the chat searchInConversation server.
func EncodeSearchInConversationRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chat.SearchInConversationPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat", "searchInConversation", "*chat.SearchInConversationPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		values.Add("searchInput", p.SearchInput)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSearchInConversationResponse returns a decoder for responses returned
// by the chat searchInConversation endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeSearchInConversationResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body SearchInConversationResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat", "searchInConversation", err)
			}
			err = ValidateSearchInConversationResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat", "searchInConversation", err)
			}
			res := NewSearchInConversationChatHistoryResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat", "searchInConversation", resp.StatusCode, string(body))
		}
	}
}

// BuildGetConversationsListRequest instantiates a HTTP request object with
// method and path set to call the "chat" service "getConversationsList"
// endpoint
func (c *Client) BuildGetConversationsListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetConversationsListChatPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat", "getConversationsList", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetConversationsListRequest returns an encoder for requests sent to
// the chat getConversationsList server.
func EncodeGetConversationsListRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chat.GetConversationsListPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat", "getConversationsList", "*chat.GetConversationsListPayload", v)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetConversationsListResponse returns a decoder for responses returned
// by the chat getConversationsList endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetConversationsListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetConversationsListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat", "getConversationsList", err)
			}
			err = ValidateGetConversationsListResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat", "getConversationsList", err)
			}
			res := NewGetConversationsListConversationsListResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat", "getConversationsList", resp.StatusCode, string(body))
		}
	}
}

// BuildAddConversationRequest instantiates a HTTP request object with method
// and path set to call the "chat" service "addConversation" endpoint
func (c *Client) BuildAddConversationRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddConversationChatPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("chat", "addConversation", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddConversationRequest returns an encoder for requests sent to the
// chat addConversation server.
func EncodeAddConversationRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*chat.AddConversationPayload)
		if !ok {
			return goahttp.ErrInvalidType("chat", "addConversation", "*chat.AddConversationPayload", v)
		}
		body := NewAddConversationRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("chat", "addConversation", err)
		}
		return nil
	}
}

// DecodeAddConversationResponse returns a decoder for responses returned by
// the chat addConversation endpoint. restoreBody controls whether the response
// body should be restored after having been read.
func DecodeAddConversationResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AddConversationResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("chat", "addConversation", err)
			}
			err = ValidateAddConversationResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("chat", "addConversation", err)
			}
			res := NewAddConversationOperationStatusResponseOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("chat", "addConversation", resp.StatusCode, string(body))
		}
	}
}

// unmarshalConversationMessageResponseBodyToChatConversationMessage builds a
// value of type *chat.ConversationMessage from a value of type
// *ConversationMessageResponseBody.
func unmarshalConversationMessageResponseBodyToChatConversationMessage(v *ConversationMessageResponseBody) *chat.ConversationMessage {
	res := &chat.ConversationMessage{
		SenderID:  *v.SenderID,
		Timestamp: *v.Timestamp,
		Content:   *v.Content,
	}

	return res
}

// unmarshalConversationResponseBodyToChatConversation builds a value of type
// *chat.Conversation from a value of type *ConversationResponseBody.
func unmarshalConversationResponseBodyToChatConversation(v *ConversationResponseBody) *chat.Conversation {
	res := &chat.Conversation{
		ID:                     *v.ID,
		LastMessageSenderID:    *v.LastMessageSenderID,
		LastMessageContent:     *v.LastMessageContent,
		LastMessageDeliveredAt: *v.LastMessageDeliveredAt,
		OtherParticipantID:     *v.OtherParticipantID,
	}

	return res
}

// marshalChatFriendToFriendRequestBody builds a value of type
// *FriendRequestBody from a value of type *chat.Friend.
func marshalChatFriendToFriendRequestBody(v *chat.Friend) *FriendRequestBody {
	res := &FriendRequestBody{
		ID:        v.ID,
		Email:     v.Email,
		FirstName: v.FirstName,
		LastName:  v.LastName,
	}

	return res
}

// marshalFriendRequestBodyToChatFriend builds a value of type *chat.Friend
// from a value of type *FriendRequestBody.
func marshalFriendRequestBodyToChatFriend(v *FriendRequestBody) *chat.Friend {
	res := &chat.Friend{
		ID:        v.ID,
		Email:     v.Email,
		FirstName: v.FirstName,
		LastName:  v.LastName,
	}

	return res
}
