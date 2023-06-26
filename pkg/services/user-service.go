package services

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/repositories"
	"github.com/jinzhu/gorm"
)

type UserService interface {
	GetAll() []User
	GetById(ID int64) (*User, *gorm.DB)
	Create(user User) *User
	Delete(ID int64) User
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{repository: userRepo}
}

type userService struct {
	repository repositories.UserRepository
}

func (u userService) GetAll() []User {
	users := u.repository.GetAll()
	return users
}

func (u userService) GetById(ID int64) (*User, *gorm.DB) {
	user, newDb := u.repository.GetById(ID)
	return user, newDb
}
func (u userService) Create(user User) *User {
	createdUser := u.repository.Create(&user)
	return createdUser
}

func (u userService) Delete(ID int64) User {
	deletedUser := u.repository.Delete(ID)
	return deletedUser
}
