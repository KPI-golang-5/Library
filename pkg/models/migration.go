package models

import (
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {

	data := db.AutoMigrate(&Book{}, &Author{}, &User{}, &UserFavBook{})
	if data == nil {
		panic("Error during migration")
	}
}
