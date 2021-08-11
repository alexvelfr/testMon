package testmon

import "github.com/alexvelfr/tmp/models"

type Notificator interface {
	Notify(msg string, recipients []models.Recipient)
}
