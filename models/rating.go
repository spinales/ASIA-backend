package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	StudentID    uint
	CourseID     uint
	Rating       uint
	RatingLetter string
	CourseName   string `gorm:"-:all"`
}
