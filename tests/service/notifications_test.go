package service

import (
	"testing"

	"github.com/mig8at/notifications/internal/model"
	"github.com/mig8at/notifications/internal/service/notifications"
	"github.com/stretchr/testify/assert"
)

func TestNotType(t *testing.T) {

	notificationService := notifications.Service()

	_, err := notificationService.Message("new", 1)

	assert.Error(t, err)

}

func TestStatus(t *testing.T) {
	notificationService := notifications.Service()
	serviceStatus, err := notificationService.Message(model.Status, 1)
	assert.NoError(t, err)

	err = serviceStatus.Send("Hola")
	assert.NoError(t, err)

	err = serviceStatus.Send("Hola 2")
	assert.NoError(t, err)

	err = serviceStatus.Send("Hola 3")
	assert.Error(t, err)

}
