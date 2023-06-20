package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func GetAllUsers() []User {
	var users []User
	newDb := db
	newDb.Find(&users)
	return users
}

func GetUserById(ID int64) (*User, *gorm.DB) {
	var getUser User
	newDb := db.Where("ID=?", ID).Find(&getUser)
	return &getUser, newDb
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

func DeleteUser(ID int64) User {
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}
