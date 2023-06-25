package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name            string `gorm:"column:name" json:"name"`
	Genre           string `gorm:"column:genre" json:"genre"`
	AuthorID        uint   `gorm:"column:author_id;foreignkey:AuthorID" json:"author_id"`
	PublicationYear int    `gorm:"column:publication_year" json:"publication_year"`
}
