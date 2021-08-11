package testmon

import "github.com/alexvelfr/tmp/models"

type Usecase interface {
	UpdateBlock(name string) error
	CheckReglamentBlocks()
	Notify(block models.Block)
}
