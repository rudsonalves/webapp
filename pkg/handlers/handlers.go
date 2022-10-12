package handlers

import (
	"net/http"
)

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request) {
	RenderTemplete(write, "home.page.html")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request) {
	RenderTemplete(write, "about.page.html")
}
