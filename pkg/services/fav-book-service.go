package services

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/repositories"
	"github.com/jinzhu/gorm"
)

type UserFavBookService interface {
	GetAll() []UserFavBook
	GetAllByUserId(userID int64) []UserFavBook
	GetById(ID int64) (*UserFavBook, *gorm.DB)
	Create(userFavBook UserFavBook) *UserFavBook
	Delete(ID int64) UserFavBook
}

func NewUserFavBookService(userFavBookRepo repositories.UserFavBookRepository) UserFavBookService {
	return &userFavBookService{repository: userFavBookRepo}
}

type userFavBookService struct {
	repository repositories.UserFavBookRepository
}

func (ufb userFavBookService) GetAll() []UserFavBook {
	userFavBooks := ufb.repository.GetAll()
	return userFavBooks
}

func (ufb userFavBookService) GetAllByUserId(userID int64) []UserFavBook {
	userFavBooks := ufb.repository.GetAllByUserId(userID)
	return userFavBooks
}

func (ufb userFavBookService) GetById(ID int64) (*UserFavBook, *gorm.DB) {
	userFavBook, newDb := ufb.repository.GetById(ID)
	return userFavBook, newDb
}
func (ufb userFavBookService) Create(userFavBook UserFavBook) *UserFavBook {
	createdUserFavBook := ufb.repository.Create(&userFavBook)
	return createdUserFavBook
}

func (ufb userFavBookService) Delete(ID int64) UserFavBook {
	deletedUserFavBook := ufb.repository.Delete(ID)
	return deletedUserFavBook
}
