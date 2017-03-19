package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// Index renders the main index page for /
func (c App) Index() revel.Result {
	var greeting = "Aloha World!"
	var books = GetBooks()
	for _, b := range books {
		revel.INFO.Printf("Index: %s | %s", b.Title, b.Author)
	}
	return c.Render(greeting, books)
}
