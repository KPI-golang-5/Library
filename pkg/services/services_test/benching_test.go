package services

import (
	"testing"

	"github.com/KPI-golang-5/Library/mock_mymodule"
	"github.com/KPI-golang-5/Library/pkg/models"
	. "github.com/KPI-golang-5/Library/pkg/services"
	"github.com/golang/mock/gomock"
)

func BenchmarkAddBook(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	bookRepository := mock_mymodule.NewMockBookRepository(ctrl)
	bookService := NewBookService(bookRepository)

	for i := 0; i < b.N; i++ {
		bookRepository.EXPECT().Create(&models.Book{Name: "testName1", Genre: "testGenre1", AuthorID: 1, PublicationYear: 1111}).Return(nil)
		bookService.Create(models.Book{Name: "testName1", Genre: "testGenre1", AuthorID: 1, PublicationYear: 1111})
	}
}

func BenchmarkGetAllBooks(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()
	bookRepository := mock_mymodule.NewMockBookRepository(ctrl)
	bookService := NewBookService(bookRepository)

	books := []models.Book{
		{Name: "testName1", Genre: "testGenre1", AuthorID: 1, PublicationYear: 1111},
		{Name: "testName2", Genre: "testGenre2", AuthorID: 2, PublicationYear: 2222},
		{Name: "testName3", Genre: "testGenre3", AuthorID: 3, PublicationYear: 3333},
		{Name: "testName4", Genre: "testGenre4", AuthorID: 4, PublicationYear: 4444},
	}

	for i := 0; i < b.N; i++ {
		bookRepository.EXPECT().GetAll("", "", "").Return(books)
		bookService.GetAll("", "", "")
	}
}

func BenchmarkFavBook(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	favBookRepository := mock_mymodule.NewMockUserFavBookRepository(ctrl)
	favBookService := NewUserFavBookService(favBookRepository)

	for i := 0; i < b.N; i++ {
		favBookRepository.EXPECT().Create(&models.UserFavBook{UserID: 1, BookID: 1}).Return(nil)
		favBookService.Create(models.UserFavBook{UserID: 1, BookID: 1})
	}
}
