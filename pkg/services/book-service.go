package services

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/repositories"
	"github.com/jinzhu/gorm"
)

type BookService interface {
	GetAll(Genre string, AuthorID string, PublicationYear string) []Book
	GetById(ID int64) (*Book, *gorm.DB)
	Create(book Book) *Book
	Delete(ID int64) Book
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{repository: bookRepo}
}

type bookService struct {
	repository repositories.BookRepository
}

func (b bookService) GetAll(Genre string, AuthorID string, PublicationYear string) []Book {
	books := b.repository.GetAll(Genre, AuthorID, PublicationYear)
	return books
}

func (b bookService) GetById(ID int64) (*Book, *gorm.DB) {
	book, newDb := b.repository.GetById(ID)
	return book, newDb
}
func (b bookService) Create(book Book) *Book {
	createdBook := b.repository.Create(&book)
	return createdBook
}

func (b bookService) Delete(ID int64) Book {
	deletedBook := b.repository.Delete(ID)
	return deletedBook
}
