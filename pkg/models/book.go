package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Book struct {
	gorm.Model
	Name            string `gorm:"column:name" json:"name"`
	Genre           string `gorm:"column:genre" json:"genre"`
	Author          string `gorm:"column:author" json:"author"`
	PublicationYear int    `gorm:"column:publication_year" json:"publication_year"`
}

func GetAllBooks(Genre string, Author string, PublicationYear string) []Book {
	var books []Book
	newDb := db
	if len(Genre) == 0 && len(Author) == 0 && len(PublicationYear) == 0 {
		// All fields are empty, return all books
		newDb.Find(&books)
	} else {
		newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
		if len(Genre) > 0 {
			newDb = newDb.Where("genre = ?", Genre)
		}
		if len(Author) > 0 {
			newDb = newDb.Where("author = ?", Author)
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
	newDb := db.Where("ID=?", ID).Find(&getBook)
	return &getBook, newDb
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
