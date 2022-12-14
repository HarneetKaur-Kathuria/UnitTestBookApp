package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBooks(t *testing.T) {

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		t.Fatal(err)
	}
	
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	// expected :=
	test := []Book{
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
	fmt.Println(test)
	fmt.Println(rr.Body.String())
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
