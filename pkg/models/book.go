package models

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name            string `gorm:"column:name" json:"name"`
	Genre           string `gorm:"column:genre" json:"genre"`
	AuthorID        uint   `gorm:"column:author_id;foreignkey:AuthorID" json:"author_id"`
	PublicationYear int    `gorm:"column:publication_year" json:"publication_year"`
}

func GetAllBooks(Genre string, AuthorID string, PublicationYear string) []Book {
	var books []Book
	newDb := Db
	if len(Genre) == 0 && len(AuthorID) == 0 && len(PublicationYear) == 0 {
		// All fields are empty, return all books
		newDb.Find(&books)
	} else {
		newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
		if len(Genre) > 0 {
			newDb = newDb.Where("genre = ?", Genre)
		}
		if len(AuthorID) > 0 {
			author, err := strconv.ParseInt(AuthorID, 0, 0)
			if err != nil {
				fmt.Println("error while parsing")
			}
			newDb = newDb.Where("author = ?", author)
		}
		if len(PublicationYear) > 0 {
			year, err := strconv.ParseInt(PublicationYear, 0, 0)
			if err != nil {
				fmt.Println("error while parsing")
			}
			newDb = newDb.Where("publication_year = ?", year)
		}
		newDb.Find(&books)
	}
	return books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	newDb := Db.Where("ID=?", ID).Find(&getBook)
	return &getBook, newDb
}

func (b *Book) CreateBook() *Book {
	Db.NewRecord(b)
	Db.Create(&b)
	return b
}

func DeleteBook(ID int64) Book {
	var book Book
	Db.Where("ID=?", ID).Delete(book)
	return book
}
