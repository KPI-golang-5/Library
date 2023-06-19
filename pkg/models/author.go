package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	FullName string    `gorm:"column:full_name" json:"full_name"`
	Country  string    `gorm:"column:country" json:"country"`
	Birth    time.Time `gorm:"column:birth" json:"birth"`
}

func GetAllAuthors(Country string) []Author {
	var authors []Author
	newDb := db
	newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
	if len(Country) > 0 {
		newDb = newDb.Where("country = ?", Country)
	}
	newDb.Find(&authors)
	return authors
}

func GetAuthorById(ID int64) (*Author, *gorm.DB) {
	var getAuthor Author
	newDb := db.Where("ID=?", ID).Find(&getAuthor)
	return &getAuthor, newDb
}

func (a *Author) CreateAuthor() *Author {
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func DeleteAuthor(ID int64) Author {
	var author Author
	db.Where("ID=?", ID).Delete(author)
	return author
}
