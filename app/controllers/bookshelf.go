package controllers

import (
	"database/sql"
	"fmt"

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
	b.Params.Bind(&book.CoverImg, "book.CoverImg")
	revel.INFO.Printf("Attempting to uploading Book titled: %s | %s", book.Title, book.Author)
	// TODO: This is empty
	revel.INFO.Printf("   Have Cover Image = %t", len(book.CoverImg) > 0)
	// revel.INFO.Printf("    PublicationDate=%s | Have CoverImg=%t", book.PublicationDate, len(book.CoverImg) > 0)
	book.Validate(b.Validation)
	if !b.Validation.HasErrors() {
		var err = Transact(m.BooksDB, func(tx *sql.Tx) error {
			uploadPK, err := tx.Prepare(`INSERT INTO books (title, author) VALUES (?, ?)`)
			if err != nil {
				revel.ERROR.Printf("Failed to create tx.Stmt: %s", err)
				return err
			}
			_, err = uploadPK.Exec(book.Title, book.Author)
			if err != nil {
				revel.ERROR.Printf("Failed to execute tx.Stmt: %s", err)
				return err
			}

			if len(book.CoverImg) > 0 {
				uploadImg, err := tx.Prepare(`UPDATE books SET cover_img=? WHERE title=? AND author=?`)
				if err != nil {
					revel.ERROR.Printf("Failed to create tx.Stmt: %s", err)
					return err
				}
				_, err = uploadImg.Exec(book.CoverImg, book.Title, book.Author)
				if err != nil {
					revel.ERROR.Printf("Failed to execute tx.Stmt: %s", err)
				}
			} else {
				var err = fmt.Errorf("Book cover image is of len 0")
				revel.ERROR.Printf("%s", err)
				return err
			}

			//TODO: Insert pub date
			// if len(PublicationDate) > 0 {
			// 	var uploadPub, err = tx.PREPARE(`INSERT INTO books`)
			// }

			// This should always be nil at this point!
			// Maybe panic...but idk what will happen inside a database transaction
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

	for name, headers := range b.Params.Files {
		revel.WARN.Printf("%s", name)
		for _, h := range headers {
			revel.WARN.Printf("    %s", h.Filename)
		}
	}

	// TODO: Figure out how to return 200 Ok without refreshing the whole page
	return b.Redirect(routes.BookShelf.Index())
}

// GetBooks returns all books in the books database
// TODO: Update this to use a transaction
func GetBooks() []*m.Book {
	var books []*m.Book

	if rows, err := m.BooksDB.Query("SELECT title, author, cover_img FROM books"); err != nil {
		revel.ERROR.Printf("Failed to run query: %s", err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var book m.Book
			if err := rows.Scan(&book.Title, &book.Author, &book.CoverImg); err != nil {
				revel.ERROR.Printf("Failed to scan row: %s", err)
			} else {
				revel.TRACE.Printf("%s | %s | coverimg=%t", book.Title, book.Author, len(book.CoverImg) > 0)
				books = append(books, &book)
			}
		}
	}
	return books
}
