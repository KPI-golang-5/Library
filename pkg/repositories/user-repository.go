package repositories

import (
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAll() []models.User
	GetById(ID int64) (*models.User, *gorm.DB)
	Create(user *models.User) *models.User
	Delete(ID int64) models.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetAll() []models.User {
	var users []models.User
	newDb := u.db
	newDb.Find(&users)
	return users
}

func (u userRepository) GetById(ID int64) (*models.User, *gorm.DB) {
	var getUser models.User
	newDb := u.db.Where("ID=?", ID).Find(&getUser)
	return &getUser, newDb
}

func (u userRepository) Create(user *models.User) *models.User {
	u.db.NewRecord(user)
	u.db.Create(&user)
	return user
}

func (u userRepository) Delete(ID int64) models.User {
	var user models.User
	u.db.Where("ID=?", ID).Delete(user)
	return user
}
