package main

import (
	"net/http"
)

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request) {
	renderTemplete(write, "home.page.html")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request) {
	renderTemplete(write, "about.page.html")
}
