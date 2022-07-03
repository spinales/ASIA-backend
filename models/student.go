package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	UserId             uint
	Career             string
	TrimesterCompleted uint8
	Pensum             string
	State              string
	QuarterlyIndex     uint
	GeneralIndex       uint
}
