package blockrepo

import (
	"github.com/alexvelfr/tmp/models"
	"github.com/alexvelfr/tmp/testmon"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

var r *repo

func NewBlockRepo(db *gorm.DB) testmon.BlockRepo {
	if r == nil {
		r = &repo{db: db}
	}
	return r
}

func (r *repo) CreateBlock(block models.Block) error {
	return r.db.Model(&block).Create(&block).Error
}

func (r *repo) UpdateBlock(block models.Block) error {
	return r.db.Save(&block).Error
}

func (r *repo) GetBlockByName(name string) (models.Block, error) {
	var res models.Block
	err := r.db.Model(&res).Where("name=?", name).First(&res).Error
	return res, err
}

func (r *repo) GetBlocks() []models.Block {
	var res []models.Block
	r.db.Model(&res).Find(&res)
	return res
}
