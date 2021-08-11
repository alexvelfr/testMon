package uc

import (
	"fmt"
	"time"

	"github.com/alexvelfr/tmp/models"
	"github.com/alexvelfr/tmp/testmon"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type uc struct {
	recipientRepo testmon.RecipientRepo
	blockRepo     testmon.BlockRepo
	notificarot   testmon.Notificator
}

var u *uc

func NewUC(recipientRepo testmon.RecipientRepo, blockRepo testmon.BlockRepo, notificarot testmon.Notificator) testmon.Usecase {
	if u == nil {
		u = &uc{
			recipientRepo: recipientRepo,
			blockRepo:     blockRepo,
			notificarot:   notificarot,
		}
	}
	return u
}

func (u *uc) UpdateBlock(name string) error {
	block, err := u.blockRepo.GetBlockByName(name)
	if err == gorm.ErrRecordNotFound {
		u.blockRepo.CreateBlock(models.Block{
			Name:          name,
			InReglament:   true,
			LastUpdate:    time.Now(),
			ReglamentTime: viper.GetDuration("app.default_duration"),
		})
		return nil
	}
	if err != nil {
		return err
	}
	if !block.InReglament {
		block.InReglament = true
		u.Notify(block)
	}
	block.LastUpdate = time.Now()
	fmt.Printf("Block %s has been updated\n", block.Name)
	return u.blockRepo.UpdateBlock(block)
}

func (u *uc) CheckReglamentBlocks() {
	blocks := u.blockRepo.GetBlocks()
	for _, b := range blocks {
		if b.InReglament && time.Since(b.LastUpdate) > b.ReglamentTime {
			b.InReglament = false
			u.Notify(b)
			u.blockRepo.UpdateBlock(b)
		}
	}
}

func (u *uc) Notify(block models.Block) {
	r := u.recipientRepo.GetRecipients()
	var msg string
	msg = fmt.Sprintf("Block %s out of reglament", block.Name)

	if block.InReglament {
		msg = fmt.Sprintf("Block %s return in reglament", block.Name)
	}
	u.notificarot.Notify(msg, r)
}
