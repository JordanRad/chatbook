// Code generated by goa v3.11.3, DO NOT EDIT.
//
// info client
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package info

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "info" service client.
type Client struct {
	GetInfoEndpoint goa.Endpoint
}

// NewClient initializes a "info" service client given the endpoints.
func NewClient(getInfo goa.Endpoint) *Client {
	return &Client{
		GetInfoEndpoint: getInfo,
	}
}

// GetInfo calls the "getInfo" endpoint of the "info" service.
func (c *Client) GetInfo(ctx context.Context) (res *OperationStatusResponse, err error) {
	var ires any
	ires, err = c.GetInfoEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*OperationStatusResponse), nil
}
