package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Selection struct {
	gorm.Model
	Code        uuid.UUID
	TrimesterID uint
	Year        string
	StudentID   uint
}
