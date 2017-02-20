package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var greeting = "Aloha World!"
	return c.Render(greeting)
}
