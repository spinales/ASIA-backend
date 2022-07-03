package models

import "gorm.io/gorm"

type Section struct {
	gorm.Model
	TeacherID uint
	CourseID  uint
	Code      string
	Schedule  string
	Quota     uint
}
