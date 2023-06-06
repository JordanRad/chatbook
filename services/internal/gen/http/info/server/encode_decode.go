// Code generated by goa v3.11.3, DO NOT EDIT.
//
// info HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package server

import (
	"context"
	"net/http"

	info "github.com/JordanRad/chatbook/services/internal/gen/info"
	goahttp "goa.design/goa/v3/http"
)

// EncodeGetInfoResponse returns an encoder for responses returned by the info
// getInfo endpoint.
func EncodeGetInfoResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*info.OperationStatusResponse)
		enc := encoder(ctx, w)
		body := NewGetInfoResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
