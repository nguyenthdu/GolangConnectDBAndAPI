package service

import (
	"connect_mysql_gorm/database"
	"connect_mysql_gorm/model"
)

type UserService interface {
	GetAllUsers() ([]model.User, error)
	GetUserByID(id uint) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type UserServiceImpl struct{}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

func (s *UserServiceImpl) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error

}

func (s *UserServiceImpl) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}

func (s *UserServiceImpl) CreateUser(user *model.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func (s *UserServiceImpl) UpdateUser(user *model.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func (s *UserServiceImpl) DeleteUser(id uint) error {
	result := database.DB.Delete(&model.User{}, id)
	return result.Error
}
