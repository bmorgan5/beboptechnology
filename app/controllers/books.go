package controllers

import (
	m "github.com/bmorgan5/beboptechnology/app/models"
	"github.com/revel/revel"
)

type BookController struct {
	*revel.Controller
}

func (c BookController) HandleBookUpload(title string) revel.Result {
	revel.INFO.Printf("Uploading Book titled: %s", title)
	return c.RenderText(title)
}

// GetBooks returns all books in the books database
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
				revel.INFO.Printf("%s | %s", book.Title, book.Author)
			}
			books = append(books, &book)
		}
	}
	return books
}
