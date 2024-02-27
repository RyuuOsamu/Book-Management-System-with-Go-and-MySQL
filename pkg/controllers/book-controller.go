package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RyuuOsamu/Book-Management-System-with-Go-and-MySQL/pkg/modules"
	"github.com/RyuuOsamu/Book-Management-System-with-Go-and-MySQL/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook modules.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	NewBooks := modules.GetAllBooks()

	res, _ := json.Marshal(NewBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}

	bookDetails, _ := modules.GetBookByID(ID)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &modules.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()

	res, _ := json.Marshal(b)
	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["bookId"]

	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}

	b := modules.DeleteBook(ID)

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}

	var updateBook = &modules.Book{}
	utils.ParseBody(r, updateBook)

	bookDetails, db := modules.GetBookByID(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
