package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KPI-golang-5/Library/pkg/models"
	. "github.com/KPI-golang-5/Library/pkg/services"
	"github.com/KPI-golang-5/Library/pkg/utils"
	"github.com/gorilla/mux"
)

type UserFavBookController struct {
	service UserFavBookService
}

func RegisterUserFavBookController(userFavBookService UserFavBookService) *UserFavBookController {
	c := &UserFavBookController{
		service: userFavBookService,
	}
	return c
}

func (c UserFavBookController) GetAllFavBooks(w http.ResponseWriter, r *http.Request) {
	favBooks := c.service.GetAll()
	res, _ := json.Marshal(favBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c UserFavBookController) GetFavBooksByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favBooks := c.service.GetAllByUserId(id)
	res, _ := json.Marshal(favBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c UserFavBookController) CreateFavBook(w http.ResponseWriter, r *http.Request) {
	createFavBook := &models.UserFavBook{}
	utils.ParseBody(r, createFavBook)
	b := c.service.Create(*createFavBook)
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c UserFavBookController) DeleteFavBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favBookId := vars["favBookId"]
	id, err := strconv.ParseInt(favBookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := c.service.Delete(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c UserFavBookController) UpdateFavBook(w http.ResponseWriter, r *http.Request) {
	var updateFavBook = &models.UserFavBook{}
	utils.ParseBody(r, updateFavBook)
	vars := mux.Vars(r)
	favBookId := vars["favBookId"]

	id, err := strconv.ParseInt(favBookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favBookRecord, db := c.service.GetById(id)
	if updateFavBook.UserID > 0 {
		favBookRecord.UserID = updateFavBook.UserID
	}
	if updateFavBook.BookID > 0 {
		favBookRecord.BookID = updateFavBook.BookID
	}

	db.Save(&favBookRecord)
	res, _ := json.Marshal(favBookRecord)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
