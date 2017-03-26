package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

// Index renders the main index page for /
func (c App) Index() revel.Result {
	return c.Render()
}
