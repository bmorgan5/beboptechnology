package controllers

import "github.com/revel/revel"

type BookController struct {
	*revel.Controller
}

type Book struct {
	Title string
}

func (c BookController) HandleBookUpload(title string) revel.Result {
	revel.INFO.Printf("Uploading Book titled: %s", title)
	return c.RenderText(title)
}

//func (c App) Index() revel.Result {
//	var greeting = "Aloha World!"
//	return c.Render(greeting)
//}
