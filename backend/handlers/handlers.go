package handlers

import (
	"book-manager/database"
	"book-manager/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


func ValidateBook(book models.Book) error {
    validate := validator.New()
    err := validate.Struct(book)
    if err != nil {
        var errorMessages []string
        for _, err := range err.(validator.ValidationErrors) {
            var errorMessage string
            switch err.Tag() {
            case "required":
                errorMessage = fmt.Sprintf("%s is required", err.Field())
            case "min":
                errorMessage = fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param())
            default:
                errorMessage = fmt.Sprintf("%s is not valid", err.Field())
            }
            errorMessages = append(errorMessages, errorMessage)
        }
        errMsg := strings.Join(errorMessages, ", ")
        log.Printf("Validation error: %s", errMsg)
        return fmt.Errorf(errMsg)
    }
    return nil
}




type ErrorResponse struct {
    Code    int    `json:"code"`    // HTTP status code
    Message string `json:"message"` // Error message
}


// GetBooks godoc
// @Summary Get list of books
// @Description Get details of all books available
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} ErrorResponse "Error retrieving books"
// @Router /books [get]
func GetBooks(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request for GetBooks: %s %s", r.Method, r.URL.Path)
    var books []models.Book
    if err := database.DB.Find(&books).Error; err != nil {
        log.Printf("Error retrieving books: %v", err)
        http.Error(w, "Error retrieving books", http.StatusInternalServerError)
        return
    }
    if err := json.NewEncoder(w).Encode(books); err != nil {
        log.Printf("Error encoding books response: %v", err)
        http.Error(w, "Error processing request", http.StatusInternalServerError)
    }
}


// AddBook adds a new book to the database
// @Summary Add a new book
// @Description Add a new book with title, author, year, genre, isbn, publisher, and description
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Add Book"
// @Success 201 {object} models.Book "Book successfully added"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Error saving book"
// @Router /books [post]
func AddBook(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received request for AddBook: %s %s", r.Method, r.URL.Path)

    var tempMap map[string]interface{}
    if err := json.NewDecoder(r.Body).Decode(&tempMap); err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    var book models.Book

    if title, ok := tempMap["title"].(string); !ok {
        http.Error(w, "title must be a string", http.StatusBadRequest)
        return
    } else {
        book.Title = title
    }

    if author, ok := tempMap["author"].(string); !ok {
        http.Error(w, "author must be a string", http.StatusBadRequest)
        return
    } else {
        book.Author = author
    }

    if year, ok := tempMap["year"].(float64); !ok {
        http.Error(w, "year must be a number", http.StatusBadRequest)
        return
    } else {
        book.Year = int(year)
    }

    if genre, ok := tempMap["genre"].(string); ok {
        book.Genre = genre
    }

    if isbn, ok := tempMap["isbn"].(string); ok {
        book.ISBN = isbn
    }

    if publisher, ok := tempMap["publisher"].(string); ok {
        book.Publisher = publisher
    }

    if description, ok := tempMap["description"].(string); ok {
        book.Description = description
    }

    if err := ValidateBook(book); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Database insertion
    if err := database.DB.Create(&book).Error; err != nil {
        log.Printf("Error saving new book: %v", err)
        http.Error(w, "Error saving book", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(book); err != nil {
        log.Printf("Error encoding book response: %v", err)
        http.Error(w, "Error processing request", http.StatusInternalServerError)
    }
}


// GetBook finds a book by its ID
// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book "Book found"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Database error"
// @Router /books/{id} [get]
func GetBook(w http.ResponseWriter, r *http.Request) {
    log.Println("GetBook request received")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Printf("Invalid ID: %v", err)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var book models.Book
    result := database.DB.First(&book, id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Printf("Book not found: %d", id)
            http.Error(w, "Book not found", http.StatusNotFound)
        } else {
            log.Printf("Database error: %v", result.Error)
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        }
        return
    }

    if err := json.NewEncoder(w).Encode(book); err != nil {
        log.Printf("Error encoding book: %v", err)
    }
}




// UpdateBook updates the details of an existing book
// @Summary Update a book
// @Description Update the details of an existing book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book object that needs to be updated"
// @Success 200 {object} models.Book "Book successfully updated"
// @Failure 400 {string} string "Invalid request body or ID"
// @Failure 404 {string} string "Book not found"
// @Failure 500 {string} string "Database error"
// @Router /books/{id} [put]
func UpdateBook(w http.ResponseWriter, r *http.Request) {
    log.Println("UpdateBook request received")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Printf("Invalid ID: %v", err)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    var book models.Book
    result := database.DB.First(&book, id)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            log.Printf("Book not found for update: %d", id)
            http.Error(w, "Book not found", http.StatusNotFound)
        } else {
            log.Printf("Database error on update: %v", result.Error)
            http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        }
        return
    }

    if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
        log.Printf("Invalid request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := ValidateBook(book); err != nil {
        log.Printf("Validation error: %v", err)
        http.Error(w, fmt.Sprintf("%v", err), http.StatusBadRequest)
        return
    }

    if err := database.DB.Save(&book).Error; err != nil {
        log.Printf("Error saving book: %v", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := json.NewEncoder(w).Encode(book); err != nil {
        log.Printf("Error encoding updated book: %v", err)
    }
}



// DeleteBook deletes a book by its ID
// @Summary Delete a book
// @Description Delete a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 204 "Book successfully deleted"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "No book found to delete"
// @Failure 500 {string} string "Error deleting book"
// @Router /books/{id} [delete]
func DeleteBook(w http.ResponseWriter, r *http.Request) {
    log.Println("DeleteBook request received")
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        log.Printf("Invalid ID for delete: %v", err)
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    result := database.DB.Delete(&models.Book{}, id)
    if result.Error != nil {
        log.Printf("Error deleting book: %v", result.Error)
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    if result.RowsAffected == 0 {
        log.Printf("No book found to delete with ID: %d", id)
        http.Error(w, "No book found to delete", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
    log.Printf("Book deleted successfully: %d", id)
}
