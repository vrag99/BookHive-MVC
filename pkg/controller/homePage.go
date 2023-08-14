package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t := views.Mode("home")
	t.Execute(w, nil)
}
