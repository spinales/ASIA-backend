package cockroach

import (
	"github.com/spinales/ASIA-backend/models"
	"gorm.io/gorm"
)

type StudentService struct {
	DB *gorm.DB
}

func (s *StudentService) AddStudent(student *models.Student) (*models.Student, error) {
	s.DB.Create(student)
	return student, nil
}

func (s *StudentService) Student(id uint) (*models.Student, error) {
	var stu models.Student
	s.DB.First(&stu, id)
	return &stu, nil
}

func (s *StudentService) Students() (*[]models.Student, error) {
	var stus []models.Student
	s.DB.Find(&stus)
	return &stus, nil
}

func (s *StudentService) DeleteStudent(id uint) error {
	s.DB.Delete(&models.Student{}, id)
	return nil
}

func (s *StudentService) UpdateStudent(student *models.Student, id uint) (*models.Student, error) {
	var c models.Student
	c.ID = id
	s.DB.Model(&c).Updates(student)
	return student, nil
}

func (s *StudentService) Ranking() (*[]models.Student, error) {
	var stus []models.Student
	s.DB.Order("general_index desc").Limit(10).Find(&stus)
	return &stus, nil
}

func (s *StudentService) StudentByUserID(userID uint) (*models.Student, error) {
	var stu models.Student
	s.DB.Where(&models.Student{UserId: userID}).First(&stu)
	return &stu, nil
}
