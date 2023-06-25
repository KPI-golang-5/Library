package repositories

import (
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/jinzhu/gorm"
)

type AuthorRepository interface {
	GetAll(country string) []models.Author
	GetById(ID int64) (*models.Author, *gorm.DB)
	Create(author *models.Author) *models.Author
	Delete(ID int64) *models.Author
}

func NewAuthorRepository(db *gorm.DB) AuthorRepository {
	return &authorRepository{db: db}
}

type authorRepository struct {
	db *gorm.DB
}

func (a authorRepository) GetAll(Country string) []models.Author {
	var authors []models.Author
	newDb := a.db
	newDb = newDb.Where("1 = 1") // Initial condition to add additional conditions
	if len(Country) > 0 {
		newDb = newDb.Where("country = ?", Country)
	}
	newDb.Find(&authors)
	return authors
}

func (a authorRepository) GetById(ID int64) (*models.Author, *gorm.DB) {
	var getAuthor models.Author
	newDb := a.db.Where("ID=?", ID).Find(&getAuthor)
	return &getAuthor, newDb
}

func (a authorRepository) Create(author *models.Author) *models.Author {
	a.db.NewRecord(author)
	a.db.Create(&author)
	return author
}

func (a authorRepository) Delete(ID int64) *models.Author {
	var author *models.Author
	a.db.Where("ID=?", ID).Delete(author)
	return author
}
