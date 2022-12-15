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

var books = []Book{
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
}

func main() {

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

// func HandlerRouting() {

// 	fmt.Println("Welcome to the DataBase Connectivity")
// 	r := mux.NewRouter()

// 	//endpoints
// 	r.HandleFunc("/", serveHome).Methods("GET")
// 	// r.HandleFunc("/book", CreateBook).Methods("POST")
// 	r.HandleFunc("/books", GetAllBooks).Methods("GET")
// 	// r.HandleFunc("/book/{id}", GetBookById).Methods("GET")
// 	r.HandleFunc("/books", DeleteAllBooks).Methods("DELETE")

// 	r.HandleFunc("/book/{id}", DeleteBookById).Methods("DELETE")
// 	// r.HandleFunc("/book/{id}", UpdateBookById).Methods("PUT")

// 	log.Fatal(http.ListenAndServe(":8080", r))

// }

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Welcome to DataBase Connectivity Session</h1>"))
// }
