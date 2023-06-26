package config

import (
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {

	data := db.AutoMigrate(&models.Book{}, &models.Author{}, &models.User{}, &models.UserFavBook{})
	if data == nil {
		panic("Error during migration")
	}
}
