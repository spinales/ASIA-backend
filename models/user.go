package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Tuition       string // matricula, lenght 11, 8 numbers and 3 letters(the short version of role)
	Password      string
	Firstname     string
	Lastname      string
	Age           uint8
	InsituteEmail string
	Status        string // Pending, Active, Inactive
	Role          string
	NationalityID uint
}
