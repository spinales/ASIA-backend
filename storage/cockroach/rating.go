package cockroach

import (
	"github.com/spinales/ASIA-backend/models"
	"gorm.io/gorm"
)

type RatingService struct {
	DB *gorm.DB
}

func (s *RatingService) Ratings(studentID uint) (*[]models.Rating, error) {
	var ratings []models.Rating
	s.DB.Where(&models.Rating{StudentID: studentID}).Find(&ratings)
	for i := 0; i < len(ratings); i++ {
		var course models.Course
		s.DB.First(&course, ratings[i].CourseID)
		ratings[i].CourseName = course.Name
	}
	return &ratings, nil
}
