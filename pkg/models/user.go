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
