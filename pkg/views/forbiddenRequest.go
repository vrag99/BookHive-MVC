package views

import (
	"html/template"
)

func ForbiddenRequest() *template.Template {
	temp := template.Must(template.ParseFiles("templates/forbiddenRequest.html", "templates/partials/headers.html"))
	return temp
}