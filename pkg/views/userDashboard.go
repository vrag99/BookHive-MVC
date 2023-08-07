package views

import (
	"html/template"
)

func UserDashboard() *template.Template {
	temp := template.Must(template.ParseFiles("templates/userDashboard.html", "templates/partials/headers.html"))
	return temp
}