package notifiation

import (
	"context"
	"fmt"

	"github.com/JordanRad/chatbook/services/internal/gen/notification"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Compile time assertion that this service implements the generated interface
var _ notification.Service = (*Service)(nil)

func (s *Service) NotifyUserNamesUpdate(ctx context.Context, p *notification.NotifyUserNamesUpdatePayload) (err error) {
	fmt.Println("User update notification: ", p)
	return nil
}
