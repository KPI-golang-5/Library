package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Author struct {
	gorm.Model
	FullName string    `gorm:"column:full_name" json:"full_name"`
	Country  string    `gorm:"column:country" json:"country"`
	Birth    time.Time `gorm:"column:birth" json:"birth"`
}
