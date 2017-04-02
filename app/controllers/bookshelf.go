package controllers

import (
	"database/sql"

	m "github.com/bmorgan5/beboptechnology/app/models"
	"github.com/bmorgan5/beboptechnology/app/routes"
	"github.com/revel/revel"
)

// BookShelf is the generic controller for dealing with the bookshelf
type BookShelf struct {
	*revel.Controller
}

// BooksJSON returns all the books in json format
func (b BookShelf) BooksJSON() revel.Result {
	// TODO: How to check if this is ajax request? Does that even matter?
	revel.INFO.Printf("Received ajax request for bookshelf")
	var books = GetBooks()
	return b.RenderJSON(books)
}

// Index renders main Index page for bookshelf app
func (b BookShelf) Index() revel.Result {
	var books = GetBooks()
	for _, b := range books {
		revel.INFO.Printf("Index: %s | %s", b.Title, b.Author)
	}
	return b.Render(books)
}

// HandleUpload receives the book upload and saves it to the books database
// TODO: Validate book input
func (b BookShelf) HandleUpload(book m.Book) revel.Result {
	revel.INFO.Printf("Attempting to uploading Book titled: %s | %s", book.Title, book.Author)
	book.Validate(b.Validation)
	if !b.Validation.HasErrors() {
		var err = Transact(m.BooksDB, func(tx *sql.Tx) error {
			var upload, err = tx.Prepare(`INSERT INTO books (title, author) VALUES (?, ?)`)
			if err != nil {
				revel.ERROR.Printf("Failed to create tx.Stmt: %s", err)
			} else {
				_, err = upload.Exec(book.Title, book.Author)
				if err != nil {
					revel.ERROR.Printf("Failed to execute tx.Stmt: %s", err)
				}
			}
			return err
		})
		if err != nil {
			revel.ERROR.Printf("Failed to HandleUpload: %s", err)
		}
	} else {
		for k, m := range b.Validation.ErrorMap() {
			revel.WARN.Printf("Validation error: %s -> %s", k, m)
		}
		b.Validation.Keep()
	}
	return b.Redirect(routes.BookShelf.Index())
}

// GetBooks returns all books in the books database
// TODO: Update this to use a transaction
func GetBooks() []*m.Book {
	var books []*m.Book

	if rows, err := m.BooksDB.Query("SELECT title, author FROM books"); err != nil {
		revel.ERROR.Printf("Failed to run query: %s", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var book m.Book
			if err := rows.Scan(&book.Title, &book.Author); err != nil {
				revel.ERROR.Printf("Failed to scan row: %s", err)
			} else {
				revel.TRACE.Printf("%s | %s", book.Title, book.Author)
				books = append(books, &book)
			}
		}
	}
	return books
}
