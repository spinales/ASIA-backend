package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	UserId             uint
	Career             string
	TrimesterCompleted uint8
	Pensum             string
	State              string
	QuarterlyIndex     float32 `gorm:"type:numeric"`
	GeneralIndex       float32 `gorm:"type:numeric"`
}
