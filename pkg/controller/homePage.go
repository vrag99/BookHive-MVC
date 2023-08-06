package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	t := views.HomePage()
	t.Execute(w, nil)
}
