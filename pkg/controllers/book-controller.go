package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/UnicoYal/go-bookstore/pkg/models"
	"github.com/UnicoYal/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["bookID"], 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, _ := models.GetBookById(bookID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	raw_book := &models.Book{}
	utils.ParseBody(r, raw_book)

	created_book := raw_book.CreateBook()

	res, _ := json.Marshal(created_book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["bookID"], 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	deleted_book := models.DeleteBook(bookID)
	res, _ := json.Marshal(deleted_book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, err := strconv.ParseInt(mux.Vars(r)["bookID"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	updated_book := &models.Book{}
	utils.ParseBody(r, updated_book)
	current_book, db := models.GetBookById(bookID)

	if updated_book.Name != "" {
		current_book.Name = updated_book.Name
	}
	if updated_book.Author != "" {
		current_book.Author = updated_book.Author
	}
	if updated_book.Publication != "" {
		current_book.Publication = updated_book.Publication
	}

	db.Save(&current_book)

	res, _ := json.Marshal(current_book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
