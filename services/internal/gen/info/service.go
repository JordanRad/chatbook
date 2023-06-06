// Code generated by goa v3.11.3, DO NOT EDIT.
//
// info service
//
// Command:
// $ goa gen github.com/JordanRad/chatbook/services/internal/design -o
// ./internal

package info

import (
	"context"
)

// Application info
type Service interface {
	// GetInfo implements getInfo.
	GetInfo(context.Context) (res *OperationStatusResponse, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "info"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"getInfo"}

// OperationStatusResponse is the result type of the info service getInfo
// method.
type OperationStatusResponse struct {
	// Operation status
	Message string
}