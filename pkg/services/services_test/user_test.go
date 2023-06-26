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

func TestNewUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("test GetAll", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			userRepo := NewMockUserRepository(ctrl)
			userService := NewUserService(userRepo)

			users := []User{
				{FullName: "John Doe", Email: "mymail@gmail.com", Password: "12345678"},
				{FullName: "John Sparrow", Email: "mymail2@gmail.com", Password: "12345678"},
				{FullName: "Megan Fox", Email: "meganfox@gmail.com", Password: "12345678"},
			}

			userRepo.EXPECT().GetAll().Return(users)

			result := userService.GetAll()
			assert.NotNil(t, result)
			assert.Equal(t, len(result), len(users))
			assert.Equal(t, result[0].FullName, users[0].FullName)
			assert.Equal(t, result[0].Email, users[0].Email)
			assert.Equal(t, result[0].Password, users[0].Password)
			assert.Equal(t, result[1].FullName, users[1].FullName)
			assert.Equal(t, result[1].Email, users[1].Email)
			assert.Equal(t, result[1].Password, users[1].Password)
		})

		t.Run("no users found", func(t *testing.T) {
			userRepo := NewMockUserRepository(ctrl)
			userService := NewUserService(userRepo)

			users := []User{}

			userRepo.EXPECT().GetAll().Return(users)

			result := userService.GetAll()
			assert.NotNil(t, result)
			assert.Equal(t, len(result), 0)
		})
	})

	t.Run("test GetById", func(t *testing.T) {
		t.Run("successful scenario", func(t *testing.T) {
			userRepo := NewMockUserRepository(ctrl)
			userService := NewUserService(userRepo)

			user := User{
				FullName: "John Doe",
				Email:    "mymail@gmail.com",
				Password: "12345678",
			}

			userRepo.EXPECT().GetById(int64(1)).Return(&user, config.GetDB())

			result, _ := userService.GetById(int64(1))
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, user.FullName)
			assert.Equal(t, result.Email, user.Email)
			assert.Equal(t, result.Password, user.Password)
		})

		t.Run("user not found", func(t *testing.T) {
			userRepo := NewMockUserRepository(ctrl)
			userService := NewUserService(userRepo)

			user := User{}
			userRepo.EXPECT().GetById(int64(5)).Return(&user, nil)

			result, _ := userService.GetById(int64(5))
			assert.NotNil(t, result)
			assert.Equal(t, result.FullName, user.FullName)
			assert.Equal(t, result.Email, user.Email)
			assert.Equal(t, result.Password, user.Password)
		})

		t.Run("test Create", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				userRepo := NewMockUserRepository(ctrl)
				userService := NewUserService(userRepo)

				user := User{
					FullName: "John Doe",
					Email:    "mymail@gmail.com",
					Password: "12345678",
				}

				userRepo.EXPECT().Create(&user).Return(&user)

				result := userService.Create(user)
				assert.NotNil(t, result)
				assert.Equal(t, result.FullName, user.FullName)
				assert.Equal(t, result.Email, user.Email)
				assert.Equal(t, result.Password, user.Password)
			})
		})

		t.Run("test Delete", func(t *testing.T) {
			t.Run("successful scenario", func(t *testing.T) {
				userRepo := NewMockUserRepository(ctrl)
				userService := NewUserService(userRepo)

				user := User{}

				userRepo.EXPECT().Delete(int64(1)).Return(user)

				result := userService.Delete(int64(1))
				assert.NotNil(t, result)
				assert.Equal(t, result.FullName, user.FullName)
				assert.Equal(t, result.Email, user.Email)
				assert.Equal(t, result.Password, user.Password)
			})
		})
	})
}
