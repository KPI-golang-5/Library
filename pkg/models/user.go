package models

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}

func GetAllUsers(Email string, Password string) []User {
	var users []User
	newDb := db
	if len(Email) == 0 && len(Password) == 0 {
		// All fields are empty, return all users
		newDb.Find(&users)
	} else {
		newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
		if len(Email) > 0 {
			newDb = newDb.Where("email = ?", Email)
		}
		if len(Password) > 0 {
			password, err := strconv.ParseInt(Password, 0, 0)
			if err != nil {
				fmt.Println("error while parsing")
			}
			newDb = newDb.Where("password = ?", password)
		}
		newDb.Find(&users)
	}
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
