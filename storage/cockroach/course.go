package cockroach

import (
	"github.com/spinales/ASIA-backend/models"
	"gorm.io/gorm"
)

type CourseService struct {
	DB *gorm.DB
}

func (s *CourseService) AddCourse(course *models.Course) (*models.Course, error) {
	s.DB.Create(course)
	return course, nil
}

func (s *CourseService) Course(id uint) (*models.Course, error) {
	var cou models.Course
	s.DB.First(&cou, id)
	return &cou, nil
}

func (s *CourseService) Courses() (*[]models.Course, error) {
	var cou []models.Course
	s.DB.Find(&cou)
	return &cou, nil
}

func (s *CourseService) CourseByName(name string) (*[]models.Course, error) {
	var courses []models.Course
	s.DB.Where(&models.Course{Name: name}).Find(&courses)
	return &courses, nil
}

func (s *CourseService) CourseByCode(code string) (*[]models.Course, error) {
	var courses []models.Course
	s.DB.Where(&models.Course{Code: code}).Find(&courses)
	return &courses, nil
}

func (s *CourseService) DeleteCourse(id uint) error {
	s.DB.Delete(&models.Course{}, id)
	return nil
}

func (s *CourseService) UpdateCourse(course *models.Course, id uint) (*models.Course, error) {
	var c models.Course
	c.ID = id
	s.DB.Model(&c).Updates(course)
	return course, nil
}
