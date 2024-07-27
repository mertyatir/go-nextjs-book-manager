package tests

import (
	"book-manager/handlers"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
    r.HandleFunc("/books", handlers.AddBook).Methods("POST")
    r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
    r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
    return r
}



func createBookForTesting(t *testing.T) string {
    var jsonStr = []byte(`{"title":"Test Book","author":"Test Author", "year":2021}`)
    request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")

    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)


    log.Printf("Response body: %s", response.Body.String())

    var result map[string]interface{}
    err := json.Unmarshal(response.Body.Bytes(), &result)
    if err != nil {
        log.Printf("Error unmarshalling response: %v", err)
        t.Fatalf("Failed to unmarshal response: %v", err)
    }

   
    idValue, ok := result["id"]
    if !ok {
        t.Fatalf("ID field is missing in the response")
    }

    bookID := strconv.Itoa(int(idValue.(float64)))
    log.Printf("Created book with ID: %s", bookID)

    return bookID
}

func TestGetBooks(t *testing.T) {
    request, _ := http.NewRequest("GET", "/books", nil)
    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Status code differs. Expected %d. Got %d instead", http.StatusOK, status)
    }
    log.Printf("TestGetBooks: Received status code %d", response.Code)
}

func TestPostBook(t *testing.T) {
    var jsonStr = []byte(`{"title":"Test Book","author":"Test Author", "year":2021}`)
    request, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")

    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)

    if status := response.Code; status != http.StatusCreated {
        t.Errorf("Status code differs. Expected %d. Got %d instead", http.StatusCreated, status)
    }
    log.Printf("TestPostBook: Received status code %d", response.Code)
}

func TestGetBook(t *testing.T) {
    bookID := createBookForTesting(t)
    request, _ := http.NewRequest("GET", "/books/"+bookID, nil)
    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Status code differs. Expected %d. Got %d instead", http.StatusOK, status)
    }
    log.Printf("TestGetBook: Received status code %d for book ID %s", response.Code, bookID)
}

func TestUpdateBook(t *testing.T) {
    bookID := createBookForTesting(t)
    var jsonStr = []byte(`{"title":"Updated Test Book","author":"Updated Test Author", "year":2022}`)
    request, _ := http.NewRequest("PUT", "/books/"+bookID, bytes.NewBuffer(jsonStr))
    request.Header.Set("Content-Type", "application/json")

    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)

    if status := response.Code; status != http.StatusOK {
        t.Errorf("Status code differs. Expected %d. Got %d instead", http.StatusOK, status)
    }
    log.Printf("TestUpdateBook: Received status code %d for book ID %s", response.Code, bookID)
}

func TestDeleteBook(t *testing.T) {
    bookID := createBookForTesting(t)
    request, _ := http.NewRequest("DELETE", "/books/"+bookID, nil)
    response := httptest.NewRecorder()
    router := setupRouter()
    router.ServeHTTP(response, request)

    if status := response.Code; status != http.StatusNoContent {
        t.Errorf("Status code differs. Expected %d. Got %d instead", http.StatusNoContent, status)
    }
    log.Printf("TestDeleteBook: Received status code %d for book ID %s", response.Code, bookID)
}