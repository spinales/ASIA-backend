package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Code         string
	Name         string
	Career       string
	Credits      uint
	AcademicArea string
}
