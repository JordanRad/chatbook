// Code generated by goa v3.11.3, DO NOT EDIT.
//
// info HTTP server types
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package server

import (
	info "github.com/JordanRad/chatbook/services/internal/gen/info"
)

// GetInfoResponseBody is the type of the "info" service "getInfo" endpoint
// HTTP response body.
type GetInfoResponseBody struct {
	// Operation status
	Message string `form:"message" json:"message" xml:"message"`
}

// NewGetInfoResponseBody builds the HTTP response body from the result of the
// "getInfo" endpoint of the "info" service.
func NewGetInfoResponseBody(res *info.OperationStatusResponse) *GetInfoResponseBody {
	body := &GetInfoResponseBody{
		Message: res.Message,
	}
	return body
}