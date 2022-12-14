package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	fmt.Println("Welcome to the DataBase Connectivity")
	r := mux.NewRouter()

	//endpoints
	r.HandleFunc("/", serveHome).Methods("GET")
	// r.HandleFunc("/book", CreateBook).Methods("POST")
	r.HandleFunc("/books", GetAllBooks).Methods("GET")
	// r.HandleFunc("/book/{id}", GetBookById).Methods("GET")
	r.HandleFunc("/books", DeleteAllBooks).Methods("DELETE")

	r.HandleFunc("/book/{id}", DeleteBookById).Methods("DELETE")
	// r.HandleFunc("/book/{id}", UpdateBookById).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to DataBase Connectivity Session</h1>"))
}
