package repositories

import (
	. "github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type AuthorRepository interface {
	GetAll(country string) []Author
	GetById(ID int64) (*Author, *gorm.DB)
	Create(author *Author) *Author
	Delete(ID int64) Author
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}

type authorRepository struct {
	db *gorm.DB
}

func (a authorRepository) GetAll(Country string) []Author {
	var authors []Author
	newDb := a.db
	newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
	if len(Country) > 0 {
		newDb = newDb.Where("country = ?", Country)
	}
	newDb.Find(&authors)
	return authors
}

func (a authorRepository) GetById(ID int64) (*Author, *gorm.DB) {
	var getAuthor Author
	newDb := a.db.Where("ID=?", ID).Find(&getAuthor)
	return &getAuthor, newDb
}

func (a authorRepository) Create(author *Author) *Author {
	a.db.NewRecord(author)
	a.db.Create(&author)
	return author
}

func (a authorRepository) Delete(ID int64) Author {
	var author Author
	a.db.Where("ID=?", ID).Delete(author)
	return author
}
