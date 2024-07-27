package models

import "time"

// Book represents a book with metadata
// @Description Book object which includes basic book information along with metadata from gorm Model
// @Property id int "The unique identifier of the book"
// @Property createdAt string "The time at which the book record was created"
// @Property updatedAt string "The time at which the book record was last updated"
// @Property deletedAt string "The time at which the book record was deleted, if applicable"
// @Property title string "The title of the book, required, minimum 2 characters"
// @Property author string "The author of the book, required, minimum 2 characters"
// @Property year int "The publication year of the book, required"
// @Property genre string "The genre of the book"
// @Property isbn string "The International Standard Book Number of the book"
// @Property publisher string "The publisher of the book"
// @Property description string "A brief description of the book"
type Book struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    CreatedAt   time.Time `json:"createdAt"`
    UpdatedAt   time.Time `json:"updatedAt"`
    DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`
    Title       string    `json:"title" validate:"required,min=2"`
    Author      string    `json:"author" validate:"required,min=2"`
    Year        int       `json:"year" validate:"required"`
    Genre       string    `json:"genre,omitempty"`
    ISBN        string    `json:"isbn,omitempty"`
    Publisher   string    `json:"publisher,omitempty"`
    Description string    `json:"description,omitempty"`
}