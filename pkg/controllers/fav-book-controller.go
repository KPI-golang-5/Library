package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KPI-golang-5/Library/pkg/models"
	"github.com/KPI-golang-5/Library/pkg/utils"
	"github.com/gorilla/mux"
)

func GetAllFavBooks(w http.ResponseWriter, r *http.Request) {
	favBooks := models.GetAllFavBooks()
	res, _ := json.Marshal(favBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFavBooksByUserID(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	id, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favBooks := models.GetFavBooksByUserID(id)
	res, _ := json.Marshal(favBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFavBook(w http.ResponseWriter, r *http.Request) {
	createFavBook := &models.UserFavBook{}
	utils.ParseBody(r, createFavBook)
	b := createFavBook.CreateFavBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFavBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	favBookId := vars["favBookId"]
	id, err := strconv.ParseInt(favBookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteFavBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateFavBook(w http.ResponseWriter, r *http.Request) {
	var updateFavBook = &models.UserFavBook{}
	utils.ParseBody(r, updateFavBook)
	vars := mux.Vars(r)
	favBookId := vars["favBookId"]

	id, err := strconv.ParseInt(favBookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	favBookRecord, db := models.GetFavBookRecord(id)
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
