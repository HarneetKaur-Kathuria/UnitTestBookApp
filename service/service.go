package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"cognologix.com/books"
	"github.com/gorilla/mux"
)


func GetAllBooks(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	json.NewEncoder(w).Encode(books.books)
}

func DeleteAllBooks(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	books.books = nil
	json.NewEncoder(w).Encode("Delete All Books")
}

func GetBookById(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	for _, book := range books.books {

		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode("ID Not Found")

}

func DeleteBookById(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")

	params := mux.Vars(request)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	for index, book := range books.books {

		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			books.books = append(books.books[:index], books.books[:index+1]...)
			json.NewEncoder(w).Encode("Book is Deleted")

			return
		}
	}

	json.NewEncoder(w).Encode("ID Not Found")

}