package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

	// fmt.Println(books)
	// Check the response body is what we expect.
	expected, err := json.Marshal(books)
	if err != nil {
		fmt.Println(err)
	}
	expectedstr := string(expected)

	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedstr) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedstr)
	}
}

func TestGetBookById(t *testing.T) {

	req, err := http.NewRequest("GET", "/book/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBookById)
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
	
	if strings.TrimSpace(rr.Body.String()) != strings.TrimSpace(expectedstr) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedstr)
	}
}
