package database

import (
	"book-manager/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
    var err error
    DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }

    DB.AutoMigrate(&models.Book{})
}