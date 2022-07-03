package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Tuition       string // matricula, lenght 11, 8 numbers and 3 letters(the short version of role)
	Firstname     string
	Lastname      string
	Age           uint8
	InsituteEmail string
	Role          string
	NationalityID uint
}
