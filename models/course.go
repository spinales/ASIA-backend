package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name         string
	Career       string
	Credits      uint
	AcademicArea string
}
