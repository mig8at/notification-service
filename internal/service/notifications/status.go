package notifications

import (
	"fmt"

	"github.com/mig8at/notifications/internal/model"
)

type StatusNotification struct {
	getUser func() (*model.User, error)
}

func (s *StatusNotification) Send(message string) error {
	user, err := s.getUser()
	fmt.Println("user:", user, err)
	return err
}
