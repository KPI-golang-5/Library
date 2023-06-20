package models

import (
	"github.com/jinzhu/gorm"
)

type UserFavBook struct {
	gorm.Model
	UserID uint `gorm:"column:user_id;foreignkey:UserID" json:"user_id"`
	BookID uint `gorm:"column:book_id;foreignkey:BookID" json:"book_id"`
}

func GetAllFavBooks() []UserFavBook {
	var favBooks []UserFavBook
	newDb := db
	newDb.Find(&favBooks)
	return favBooks
}

func GetFavBooksByUserID(userID int64) []UserFavBook {
	var favBooks []UserFavBook
	newDb := db
	newDb.Where("user_id = ?", userID).Find(&favBooks)
	return favBooks
}

func GetFavBookRecord(ID int64) (*UserFavBook, *gorm.DB) {
	var getFavBookRecord UserFavBook
	newDb := db.Where("ID=?", ID).Find(&getFavBookRecord)
	return &getFavBookRecord, newDb
}

func (b *UserFavBook) CreateFavBook() *UserFavBook {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func DeleteFavBook(favBookID int64) UserFavBook {
	var favBook UserFavBook
	db.Where("ID=?", favBookID).Delete(favBook)
	return favBook
}
