package models

import (
	"github.com/KPI-golang-5/Library/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	Migrate(db)
}
