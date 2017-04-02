package controllers

import (
	"github.com/revel/revel"
)

// App is the main website controller
// TODO: Rename this to Home or something
type App struct {
	*revel.Controller
}

// Index renders the main index page for /
func (c App) Index() revel.Result {
	return c.Render()
}
