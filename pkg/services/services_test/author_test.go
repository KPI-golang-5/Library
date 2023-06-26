package services_test

import (
	. "github.com/KPI-golang-5/Library/mock_mymodule"
	"github.com/KPI-golang-5/Library/pkg/config"
	. "github.com/KPI-golang-5/Library/pkg/models"
	. "github.com/KPI-golang-5/Library/pkg/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewAuthorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test GetAll", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			authors := []Author{
				{FullName: "John Doe", Country: "Sweden", Birth: time.Now()},
				{FullName: "Taras Shevchenko", Country: "Ukraine", Birth: time.Now()},
				{FullName: "Oskar Wilde", Country: "Sweden", Birth: time.Now()},
			}

			authorRepo.EXPECT().GetAll("").Return(authors)

			result := authorService.GetAll("")
			assert.NotNil(t, result)
			assert.Equal(t, len(result), len(authors))
		})

		t.Run("no authors found", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			authorRepo.EXPECT().GetAll("").Return(nil)

			result := authorService.GetAll("")
			assert.Equal(t, len(result), 0)
		})
	})

	t.Run("test GetById", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			author := Author{
				FullName: "John Doe",
				Country:  "Sweden",
				Birth:    time.Now(),
			}

			authorRepo.EXPECT().GetById(int64(1)).Return(&author, config.GetDB())

			result, _ := authorService.GetById(int64(1))
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, author.FullName)
			assert.Equal(t, result.Country, author.Country)
			assert.Equal(t, result.Birth, author.Birth)
		})

		t.Run("author not found", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			author := Author{}
			authorRepo.EXPECT().GetById(int64(5)).Return(&author, nil)

			result, _ := authorService.GetById(int64(5))
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, author.FullName)
			assert.Equal(t, result.Country, author.Country)
			assert.Equal(t, result.Birth, author.Birth)
		})
	})

	t.Run("test Create", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			author := Author{
				FullName: "John Doe",
				Country:  "Sweden",
				Birth:    time.Now(),
			}

			authorRepo.EXPECT().Create(&author).Return(&author)

			result := authorService.Create(author)
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, author.FullName)
			assert.Equal(t, result.Country, author.Country)
			assert.Equal(t, result.Birth, author.Birth)
		})
	})

	t.Run("test Delete", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			authorRepo := NewMockAuthorRepository(ctrl)
			authorService := NewAuthorService(authorRepo)

			author := Author{}

			authorRepo.EXPECT().Delete(int64(1)).Return(author)

			result := authorService.Delete(int64(1))
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, author.FullName)
			assert.Equal(t, result.Country, author.Country)
			assert.Equal(t, result.Birth, author.Birth)
		})
	})
}
