package views

import (
	"html/template"
)

func NotFound() *template.Template {
	temp := template.Must(template.ParseFiles("templates/notFound.html", "templates/partials/headers.html"))
	return temp
}