package testmon

import "github.com/alexvelfr/tmp/models"

type BlockRepo interface {
	CreateBlock(models.Block) error
	UpdateBlock(models.Block) error
	GetBlockByName(name string) (models.Block, error)
	GetBlocks() []models.Block
}

type RecipientRepo interface {
	GetRecipients() []models.Recipient
}
