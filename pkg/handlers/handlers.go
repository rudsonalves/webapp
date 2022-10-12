package handlers

import (
	"net/http"

	"github.com/rudsonalves/webapp/pkg/render"
)

// Home is the home page handler
func Home(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplete(write, "home.page.html")
}

// About is the about page handler
func About(write http.ResponseWriter, request *http.Request) {
	render.RenderTemplete(write, "about.page.html")
}
