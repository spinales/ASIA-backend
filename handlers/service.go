package handlers

import (
	"github.com/spinales/ASIA-backend/storage/cockroach"
	"gorm.io/gorm"
)

// service represent all of the dependency required for the app
type Service struct {
	UserService    *cockroach.UserService
	CourseService  *cockroach.CourseService
	StudentService *cockroach.StudentService
	RatingService  *cockroach.RatingService
}

// NewService create new service
func NewService(db *gorm.DB) *Service {
	return &Service{
		UserService:    &cockroach.UserService{DB: db},
		CourseService:  &cockroach.CourseService{DB: db},
		StudentService: &cockroach.StudentService{DB: db},
		RatingService:  &cockroach.RatingService{DB: db},
	}
}
