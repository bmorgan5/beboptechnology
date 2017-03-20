package models

import (
	"database/sql"

	"github.com/revel/revel"
)

// BooksDB is the Handle too books sqlite database
// TODO: This should move to controllers/db.go and be generic db handle
var BooksDB *sql.DB

// Book is the go structure for a book taken from the BooksDB
type Book struct {
	Title           string
	Author          string
	PublicationDate string //TODO: Make this a real date object
}

// Validates that an inserted book has all these fields
func (book Book) Validate(v *revel.Validation) {
	v.Required(book.Title).Message("A title is required")
	v.Required(book.Author).Message("An author is required")
}
