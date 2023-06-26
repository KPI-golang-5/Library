package controllers

import (
	"encoding/json"
	"fmt"
	. "github.com/KPI-golang-5/Library/pkg/services"
	"net/http"
	"strconv"
	"time"

	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/utils"
	"github.com/gorilla/mux"
)

type AuthorController struct {
	service AuthorService
}

func RegisterAuthorController(authorService AuthorService) *AuthorController {
	c := &AuthorController{
		service: authorService,
	}
	return c
}

func (c AuthorController) GetAuthors(w http.ResponseWriter, r *http.Request) {
	country := r.URL.Query().Get("country")
	authors := c.service.GetAll(country)
	res, _ := json.Marshal(authors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c AuthorController) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	id, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authorDetails, _ := c.service.GetById(id)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c AuthorController) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	createAuthor := &models.Author{}
	utils.ParseBody(r, createAuthor)
	b := c.service.Create(*createAuthor)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c AuthorController) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	id, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	author := c.service.Delete(id)
	res, _ := json.Marshal(author)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c AuthorController) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var updateAuthor = &models.Author{}
	utils.ParseBody(r, updateAuthor)
	vars := mux.Vars(r)
	authorId := vars["authorId"]
	id, err := strconv.ParseInt(authorId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	authorDetails, db := c.service.GetById(id)
	if updateAuthor.FullName != "" {
		authorDetails.FullName = updateAuthor.FullName
	}
	if updateAuthor.Country != "" {
		authorDetails.Country = updateAuthor.Country
	}
	if updateAuthor.Birth.After(time.Date(0, time.January, 1, 0, 0, 0, 0, time.UTC)) {
		authorDetails.Birth = updateAuthor.Birth
	}
	db.Save(&authorDetails)
	res, _ := json.Marshal(authorDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
