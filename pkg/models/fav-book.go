package models

import (
	"github.com/jinzhu/gorm"
)

type UserFavBook struct {
	gorm.Model
	UserID uint `gorm:"column:user_id;foreignkey:UserID" json:"user_id"`
	BookID uint `gorm:"column:book_id;foreignkey:BookID" json:"book_id"`
}
