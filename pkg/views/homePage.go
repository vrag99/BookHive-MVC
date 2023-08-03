package views

import (
	"html/template"
)

func HomePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/home.html", "templates/partials/headers.html"))
	return temp
}