package repositories

import (
	"fmt"
	"github.com/KPI-golang-5/Library/pkg/config"
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
	"strconv"
)

type BookRepository interface {
	GetAll(Genre string, AuthorID string, PublicationYear string) []Book
	GetById(ID int64) (*Book, *gorm.DB)
	Create(b *Book) *Book
	Delete(ID int64) Book
}

func NewBookRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}

type bookRepository struct {
	db *gorm.DB
}

func (b bookRepository) GetAll(Genre string, AuthorID string, PublicationYear string) []Book {
	var books []Book
	newDb := config.GetDB()
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

func (b bookRepository) GetById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	newDb := config.GetDB().Where("ID=?", ID).Find(&getBook)
	return &getBook, newDb
}

func (b bookRepository) Create(book *Book) *Book {
	config.GetDB().NewRecord(book)
	config.GetDB().Create(&book)
	return book
}

func (b bookRepository) Delete(ID int64) Book {
	var book Book
	config.GetDB().Where("ID=?", ID).Delete(book)
	return book
}
