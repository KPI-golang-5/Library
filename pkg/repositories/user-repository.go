package repositories

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	GetAll() []User
	GetById(ID int64) (*User, *gorm.DB)
	Create(user *User) *User
	Delete(ID int64) User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) GetAll() []User {
	var users []User
	newDb := u.db
	newDb.Find(&users)
	return users
}

func (u userRepository) GetById(ID int64) (*User, *gorm.DB) {
	var getUser User
	newDb := u.db.Where("ID=?", ID).Find(&getUser)
	return &getUser, newDb
}

func (u userRepository) Create(user *User) *User {
	u.db.NewRecord(user)
	u.db.Create(&user)
	return user
}

func (u userRepository) Delete(ID int64) User {
	var user User
	u.db.Where("ID=?", ID).Delete(user)
	return user
}
