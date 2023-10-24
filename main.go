package main

import (
	"github.com/mig8at/notifications/internal/model"
	"github.com/mig8at/notifications/internal/service/notifications"
)

func main() {

	notificationService := notifications.Service()

	stateMessage, err := notificationService.Message(model.Status, 1)
	if err != nil {
		panic(err)
	}

	stateMessage.Send("Hola")
	stateMessage.Send("Hola 2")
	stateMessage.Send("Hola 3")
	stateMessage.Send("Hola 4")
}
