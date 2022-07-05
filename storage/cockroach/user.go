package cockroach

import (
	"github.com/spinales/ASIA-backend/models"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

// UserByTuition search the user in db by tuition
func (s *UserService) UserByTuition(tuition string) (*models.User, error) {
	var user models.User
	s.DB.Where("tuition=?", tuition).First(&user)
	return &user, nil
}

func (s *UserService) AddUser(user *models.User) (*models.User, error) {
	s.DB.Create(user)
	return user, nil
}

func (s *UserService) User(id uint) (*models.User, error) {
	var cou models.User
	s.DB.First(&cou, id)
	return &cou, nil
}

func (s *UserService) Users() (*[]models.User, error) {
	var users []models.User
	s.DB.Find(&users)
	return &users, nil
}

func (s *UserService) UserByFirstname(name string) (*models.User, error) {
	var user models.User
	s.DB.Where(&models.User{Firstname: name}).First(&user)
	return &user, nil
}

func (s *UserService) Deleteuser(id uint) error {
	s.DB.Delete(&models.User{}, id)
	return nil
}

func (s *UserService) UpdateUser(user *models.User, id uint) (*models.User, error) {
	var c models.User
	c.ID = id
	s.DB.Model(&c).Updates(user)
	return user, nil
}
