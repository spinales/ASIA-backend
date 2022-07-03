package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	UserId       uint
	AcademicArea string
	Sections     []Section
}
