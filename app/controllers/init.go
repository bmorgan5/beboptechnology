package controllers

import "github.com/revel/revel"

// This function get's called magically by revel...idk how
func init() {
	revel.INFO.Printf("controllers: init()")
	revel.OnAppStart(InitBookDB)
}
