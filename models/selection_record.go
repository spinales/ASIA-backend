package models

import "gorm.io/gorm"

type SelectionRecord struct {
	gorm.Model
	SectionID   uint
	SelectionID uint
}
