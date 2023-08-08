package views

import (
	"html/template"
)

func AdminDashboard() *template.Template {
	temp := template.Must(template.ParseFiles("templates/adminDashboard.html", "templates/partials/headers.html"))
	return temp
}