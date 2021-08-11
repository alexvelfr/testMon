package recipientrepo

import (
	"github.com/alexvelfr/tmp/models"
	"github.com/alexvelfr/tmp/testmon"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

var r *repo

func NewBlockRepo(db *gorm.DB) testmon.RecipientRepo {
	if r == nil {
		r = &repo{db: db}
	}
	return r
}

func (r *repo) GetRecipients() []models.Recipient {
	var res []models.Recipient
	r.db.Model(&models.Recipient{}).Find(&res)
	return res
}
