package models

import (
	"database/sql"
)

// BooksDB is the Handle too books sqlite database
var BooksDB *sql.DB

// Book is the go structure for a book taken from the BooksDB
type Book struct {
	Title           string
	Author          string
	PublicationDate string //TODO: Make this a real date object
}
