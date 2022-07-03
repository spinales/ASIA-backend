package models

import (
	"time"

	"gorm.io/gorm"
)

type Trimester struct {
	gorm.Model
	Name       string
	InitDate   time.Time
	FinishDate time.Time
}
