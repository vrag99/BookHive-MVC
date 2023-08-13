package views

import (
	"html/template"
)

func InternalServerError() *template.Template {
	temp := template.Must(template.ParseFiles("templates/internalServerError.html", "templates/partials/headers.html"))
	return temp
}