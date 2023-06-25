package repositories

import (
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserFavBookRepository interface {
	GetAll() []models.UserFavBook
	GetAllByUserID(userID int64) []models.UserFavBook
	GetByID(ID int64) (*models.UserFavBook, *gorm.DB)
	Create(b *models.UserFavBook) *models.UserFavBook
	Delete(favBookID int64) models.UserFavBook
}

func NewUserFavBookRepository(db *gorm.DB) UserFavBookRepository {
	return &userFavBookRepository{db: db}
}

type userFavBookRepository struct {
	db *gorm.DB
}

func (ufb userFavBookRepository) GetAll() []models.UserFavBook {
	var favBooks []models.UserFavBook
	newDb := ufb.db
	newDb.Find(&favBooks)
	return favBooks
}

func (ufb userFavBookRepository) GetAllByUserID(userID int64) []models.UserFavBook {
	var favBooks []models.UserFavBook
	newDb := ufb.db
	newDb.Where("user_id = ?", userID).Find(&favBooks)
	return favBooks
}

func (ufb userFavBookRepository) GetByID(ID int64) (*models.UserFavBook, *gorm.DB) {
	var getFavBookRecord models.UserFavBook
	newDb := ufb.db.Where("ID=?", ID).Find(&getFavBookRecord)
	return &getFavBookRecord, newDb
}

func (ufb userFavBookRepository) Create(b *models.UserFavBook) *models.UserFavBook {
	ufb.db.NewRecord(b)
	ufb.db.Create(&b)
	return b
}

func (ufb userFavBookRepository) Delete(favBookID int64) models.UserFavBook {
	var favBook models.UserFavBook
	ufb.db.Where("ID=?", favBookID).Delete(favBook)
	return favBook
}
