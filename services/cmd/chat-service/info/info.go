package info

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JordanRad/chatbook/services/internal/gen/info"
)

type Service struct {
	logger *log.Logger
}

// Compile time assertion that this service implements the generated interface
var _ info.Service = (*Service)(nil)

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetInfo(ctx context.Context) (*info.OperationStatusResponse, error) {

	response := &info.OperationStatusResponse{
		Message: fmt.Sprintf("[%v] User Management service is up and running", time.Now().Format("2 Jan 2006 15:04")),
	}
	return response, nil
}
