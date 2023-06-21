package models

import (
	"github.com/KPI-golang-5/Library/pkg/config"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	config.Connect()
	Db = config.GetDB()
	Migrate(Db)
}
