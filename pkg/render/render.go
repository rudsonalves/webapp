package render

import (
	"log"
	"net/http"
	"text/template"
)

// RenderTemplete renders templetes using html/templetes
func RenderTemplete(write http.ResponseWriter, tmpl string) {
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
