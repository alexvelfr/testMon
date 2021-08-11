package models

import (
	"time"

	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Name          string
	ReglamentTime time.Duration
	LastUpdate    time.Time
	InReglament   bool
}
