package main

import (
	"encoding/json"
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

	// defer req.Body.Close()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBooks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// fmt.Println(books)
	// Check the response body is what we expect.
	expected, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
	}
	expectedstr := string(expected)
	if rr.Body.String() != expectedstr {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedstr)
	}
}
