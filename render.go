package main

import (
	"log"
	"net/http"
	"text/template"
)

func renderTemplete(write http.ResponseWriter, tmpl string) {
	parsedTemplete, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println(err)
		return
	}

	err = parsedTemplete.Execute(write, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
