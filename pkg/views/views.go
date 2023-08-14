package views

import "html/template"

func Mode(viewMode string) *template.Template {
	mode := map[string]*template.Template{
		"home":                getTemplate("home"),
		"adminDashboard":      getTemplate("adminDashboard"),
		"userDashboard":       getTemplate("userDashboard"),
		"login":               getTemplate("login"),
		"register":            getTemplate("register"),
		"notFound":            getTemplate("notFound"),
		"internalServerError": getTemplate("internalServerError"),
		"forbiddenRequest":    getTemplate("forbiddenRequest"),
	}

	temp := mode[viewMode]
	return temp
}

func getTemplate(viewMode string) *template.Template {
	path := "templates/" + viewMode
	return template.Must(template.ParseFiles(path + ".html", "templates/partials/headers.html"))
}
