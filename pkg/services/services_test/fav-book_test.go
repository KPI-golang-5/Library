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

func TestNewUserFavBookService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test GetAll", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			userFavBookRepo := NewMockUserFavBookRepository(ctrl)
			userFavBookService := NewUserFavBookService(userFavBookRepo)

			userFavBooks := []UserFavBook{
				{UserID: 1, BookID: 1},
				{UserID: 1, BookID: 2},
				{UserID: 2, BookID: 3},
			}

			userFavBookRepo.EXPECT().GetAll().Return(userFavBooks)

			result := userFavBookService.GetAll()
			assert.NotNil(t, result)
			assert.Equal(t, len(result), len(userFavBooks))
			assert.Equal(t, result[0].UserID, userFavBooks[0].UserID)
			assert.Equal(t, result[0].BookID, userFavBooks[0].BookID)
			assert.Equal(t, result[1].UserID, userFavBooks[1].UserID)
			assert.Equal(t, result[1].BookID, userFavBooks[1].BookID)
		})

		t.Run("no userFavBooks found", func(t *testing.T) {
			userFavBookRepo := NewMockUserFavBookRepository(ctrl)
			userFavBookService := NewUserFavBookService(userFavBookRepo)

			userFavBooks := []UserFavBook{}

			userFavBookRepo.EXPECT().GetAll().Return(userFavBooks)

			result := userFavBookService.GetAll()
			assert.NotNil(t, result)
			assert.Equal(t, len(result), 0)
		})
	})

	t.Run("test GetById", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			userFavBookRepo := NewMockUserFavBookRepository(ctrl)
			userFavBookService := NewUserFavBookService(userFavBookRepo)

			userFavBook := UserFavBook{
				UserID: 1,
				BookID: 1,
			}

			userFavBookRepo.EXPECT().GetById(int64(1)).Return(&userFavBook, config.GetDB())

			result, _ := userFavBookService.GetById(int64(1))
			assert.NotNil(t, result)
			assert.Equal(t, result.UserID, userFavBook.UserID)
			assert.Equal(t, result.BookID, userFavBook.BookID)
		})

		t.Run("userFavBook not found", func(t *testing.T) {
			userFavBookRepo := NewMockUserFavBookRepository(ctrl)
			userFavBookService := NewUserFavBookService(userFavBookRepo)

			userFavBook := UserFavBook{}
			userFavBookRepo.EXPECT().GetById(int64(5)).Return(&userFavBook, nil)

			result, _ := userFavBookService.GetById(int64(5))
			assert.NotNil(t, result)
			assert.Equal(t, result.UserID, userFavBook.UserID)
			assert.Equal(t, result.BookID, userFavBook.BookID)
		})

		t.Run("test Create", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				userFavBookRepo := NewMockUserFavBookRepository(ctrl)
				userFavBookService := NewUserFavBookService(userFavBookRepo)

				userFavBook := UserFavBook{
					UserID: 1,
					BookID: 1,
				}

				userFavBookRepo.EXPECT().Create(&userFavBook).Return(&userFavBook)

				result := userFavBookService.Create(userFavBook)
				assert.NotNil(t, result)
				assert.Equal(t, result.UserID, userFavBook.UserID)
				assert.Equal(t, result.BookID, userFavBook.BookID)
			})
		})

		t.Run("test Delete", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				userFavBookRepo := NewMockUserFavBookRepository(ctrl)
				userFavBookService := NewUserFavBookService(userFavBookRepo)

				userFavBook := UserFavBook{}

				userFavBookRepo.EXPECT().Delete(int64(1)).Return(userFavBook)

				result := userFavBookService.Delete(int64(1))
				assert.NotNil(t, result)
				assert.Equal(t, result.UserID, userFavBook.UserID)
				assert.Equal(t, result.BookID, userFavBook.BookID)
			})
		})
	})
}
