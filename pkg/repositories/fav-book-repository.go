package repositories

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserFavBookRepository interface {
	GetAll() []UserFavBook
	GetAllByUserId(userID int64) []UserFavBook
	GetById(ID int64) (*UserFavBook, *gorm.DB)
	Create(favBook *UserFavBook) *UserFavBook
	Delete(favBookID int64) UserFavBook
}

func NewUserFavBookRepository(db *gorm.DB) UserFavBookRepository {
	return &userFavBookRepository{db: db}
}

type userFavBookRepository struct {
	db *gorm.DB
}

func (ufb userFavBookRepository) GetAll() []UserFavBook {
	var favBooks []UserFavBook
	newDb := ufb.db
	newDb.Find(&favBooks)
	return favBooks
}

func (ufb userFavBookRepository) GetAllByUserId(userID int64) []UserFavBook {
	var favBooks []UserFavBook
	newDb := ufb.db
	newDb.Where("user_id = ?", userID).Find(&favBooks)
	return favBooks
}

func (ufb userFavBookRepository) GetById(ID int64) (*UserFavBook, *gorm.DB) {
	var getFavBookRecord UserFavBook
	newDb := ufb.db.Where("ID=?", ID).Find(&getFavBookRecord)
	return &getFavBookRecord, newDb
}

func (ufb userFavBookRepository) Create(b *UserFavBook) *UserFavBook {
	ufb.db.NewRecord(b)
	ufb.db.Create(&b)
	return b
}

func (ufb userFavBookRepository) Delete(favBookID int64) UserFavBook {
	var favBook UserFavBook
	ufb.db.Where("ID=?", favBookID).Delete(favBook)
	return favBook
}
