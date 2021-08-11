package notificator

import (
	"fmt"

	"github.com/alexvelfr/tmp/models"
	"github.com/alexvelfr/tmp/testmon"
)

type consoleNotificarot struct {
}

func NewNotificator() testmon.Notificator {
	return &consoleNotificarot{}
}

func (n *consoleNotificarot) Notify(msg string, recipients []models.Recipient) {
	fmt.Println(msg)
}
