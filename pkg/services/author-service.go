package services

import (
	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/repositories"
	"github.com/jinzhu/gorm"
)

type AuthorService interface {
	GetAll(country string) []models.Author
	GetById(ID int64) (*models.Author, *gorm.DB)
	Create(author models.Author) *models.Author
	Delete(ID int64) *models.Author
}

func NewAuthorService(authorRepo repositories.AuthorRepository) AuthorService {
	return &authorService{repository: authorRepo}
}

type authorService struct {
	repository repositories.AuthorRepository
}

func (a authorService) GetAll(country string) []models.Author {
	authors := a.repository.GetAll(country)
	return authors
}

func (a authorService) GetById(ID int64) (*models.Author, *gorm.DB) {
	author, newDb := a.repository.GetById(ID)
	return author, newDb
}
func (a authorService) Create(author models.Author) *models.Author {
	createdAuthor := a.repository.Create(&author)
	return createdAuthor
}

func (a authorService) Delete(ID int64) *models.Author {
	deletedAuthor := a.repository.Delete(ID)
	return deletedAuthor
}

/*func mapModelToResp(authorModel *models.Author) *Author {

	a := &Author{
		Model:    authorModel.Model,
		FullName: authorModel.FullName,
		Country:  authorModel.Country,
		Birth:    authorModel.Birth,
	}

	return a
}
*/
