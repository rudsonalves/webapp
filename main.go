package main

import (
	"net/http"
)

const portNumber = ":6060"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
