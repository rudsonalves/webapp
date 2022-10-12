package main

import (
	"net/http"

	"github.com/rudsonalves/webapp/pkg/handlers"
)

const portNumber = ":6060"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(portNumber, nil)
}
