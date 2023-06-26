package services_test

import (
	. "github.com/KPI-golang-5/Library/mock_mymodule"
	"github.com/KPI-golang-5/Library/pkg/config"
	. "github.com/KPI-golang-5/Library/pkg/models"
	. "github.com/KPI-golang-5/Library/pkg/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test GetAll", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			bookRepo := NewMockBookRepository(ctrl)
			bookService := NewBookService(bookRepo)

			books := []Book{
				{Name: "Fairy Tale", Genre: "Fantasy", AuthorID: 1, PublicationYear: 2020},
				{Name: "Fairy Tale 2", Genre: "Fantasy", AuthorID: 1, PublicationYear: 2021},
				{Name: "Fairy Tale 3", Genre: "Fantasy", AuthorID: 1, PublicationYear: 2023},
			}

			bookRepo.EXPECT().GetAll("", "", "").Return(books)

			result := bookService.GetAll("", "", "")
			assert.NotNil(t, result)
			assert.Equal(t, len(result), len(books))
			assert.Equal(t, result[0].Name, books[0].Name)
			assert.Equal(t, result[0].Genre, books[0].Genre)
			assert.Equal(t, result[0].AuthorID, books[0].AuthorID)
			assert.Equal(t, result[0].PublicationYear, books[0].PublicationYear)
			assert.Equal(t, result[1].Name, books[1].Name)
			assert.Equal(t, result[1].Genre, books[1].Genre)
			assert.Equal(t, result[1].AuthorID, books[1].AuthorID)
			assert.Equal(t, result[1].PublicationYear, books[1].PublicationYear)
		})

		t.Run("no books found", func(t *testing.T) {
			bookRepo := NewMockBookRepository(ctrl)
			bookService := NewBookService(bookRepo)

			books := []Book{}

			bookRepo.EXPECT().GetAll("Horror", "", "").Return(books)

			result := bookService.GetAll("Horror", "", "")
			assert.NotNil(t, result)
			assert.Equal(t, len(result), 0)
		})
	})

	t.Run("test GetById", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			bookRepo := NewMockBookRepository(ctrl)
			bookService := NewBookService(bookRepo)

			book := Book{
				Name:            "Fairy Tale",
				Genre:           "Fantasy",
				AuthorID:        1,
				PublicationYear: 2020,
			}

			bookRepo.EXPECT().GetById(int64(1)).Return(&book, config.GetDB())

			result, _ := bookService.GetById(int64(1))
			assert.NotNil(t, result)
			assert.Equal(t, result.Name, book.Name)
			assert.Equal(t, result.Genre, book.Genre)
			assert.Equal(t, result.AuthorID, book.AuthorID)
			assert.Equal(t, result.PublicationYear, book.PublicationYear)
		})

		t.Run("book not found", func(t *testing.T) {
			bookRepo := NewMockBookRepository(ctrl)
			bookService := NewBookService(bookRepo)

			book := Book{}
			bookRepo.EXPECT().GetById(int64(5)).Return(&book, nil)

			result, _ := bookService.GetById(int64(5))
			assert.NotNil(t, result)
			assert.Equal(t, result.Name, book.Name)
			assert.Equal(t, result.Genre, book.Genre)
			assert.Equal(t, result.AuthorID, book.AuthorID)
			assert.Equal(t, result.PublicationYear, book.PublicationYear)
		})

		t.Run("test Create", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				bookRepo := NewMockBookRepository(ctrl)
				bookService := NewBookService(bookRepo)

				book := Book{
					Name:            "Fairy Tale",
					Genre:           "Fantasy",
					AuthorID:        1,
					PublicationYear: 2020,
				}

				bookRepo.EXPECT().Create(&book).Return(&book)

				result := bookService.Create(book)
				assert.NotNil(t, result)
				assert.Equal(t, result.Name, book.Name)
				assert.Equal(t, result.Genre, book.Genre)
				assert.Equal(t, result.AuthorID, book.AuthorID)
				assert.Equal(t, result.PublicationYear, book.PublicationYear)
			})
		})

		t.Run("test Delete", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				bookRepo := NewMockBookRepository(ctrl)
				bookService := NewBookService(bookRepo)

				book := Book{}

				bookRepo.EXPECT().Delete(int64(1)).Return(book)

				result := bookService.Delete(int64(1))
				assert.NotNil(t, result)
				assert.Equal(t, result.Name, book.Name)
				assert.Equal(t, result.Genre, book.Genre)
				assert.Equal(t, result.AuthorID, book.AuthorID)
				assert.Equal(t, result.PublicationYear, book.PublicationYear)
			})
		})
	})
}
