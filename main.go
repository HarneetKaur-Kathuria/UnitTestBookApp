package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID            int    `json:"courseId"`
	Title         string `json:"bookName"`
	Author        string `json:"writer"`
	YearPublished int    `json:"published year"`
}

var books []Book

func main() {

	books = []Book{
		{
			ID:            1,
			Title:         "Introduction to Linux",
			Author:        "Douglas Adams",
			YearPublished: 1979,
		},
		{
			ID:            2,
			Title:         "The Hobbit",
			Author:        "J.R.R Tolkein",
			YearPublished: 1937,
		},
		{
			ID:            3,
			Title:         "A Tale of Two Cities",
			Author:        "Charles",
			YearPublished: 2009,
		},
		{
			ID:            4,
			Title:         "Harry Potter and the Philosophers Stone",
			Author:        "J.K. Rowling",
			YearPublished: 1997,
		},
		{
			ID:            5,
			Title:         "Les Miserables",
			Author:        "Victor Hugo",
			YearPublished: 1862,
		},
		{
			ID:            6,
			Title:         "I, Robot",
			Author:        "Isaac Asamov",
			YearPublished: 1950,
		},
		{
			ID:            7,
			Title:         "The Gods Themselves",
			Author:        "Isaac Asamov",
			YearPublished: 1973,
		},
		{
			ID:            8,
			Title:         "The Moond is a Hash Mistress",
			Author:        "Robert A. Heinlein",
			YearPublished: 1966,
		},
		{
			ID:            9,
			Title:         "On Basilisk Station",
			Author:        "David Weber",
			YearPublished: 1993,
		},
		{
			ID:            10,
			Title:         "The Android's Dream",
			Author:        "John Scalzi",
			YearPublished: 2006,
		},
	}
	HandlerRouting()
}

func GetAllBooks(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func DeleteAllBooks(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content - Type", "application/json")
	books = nil
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

	for _, book := range books {

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

	for index, book := range books {

		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			books = append(books[:index], books[:index+1]...)
			json.NewEncoder(w).Encode("Book is Deleted")

			return
		}
	}

	json.NewEncoder(w).Encode("ID Not Found")

}
