package controller

import (
	"net/http"
	"BookHive/pkg/views"
)

func HomePage(writer http.ResponseWriter, request *http.Request) {
	t := views.HomePage()
	t.Execute(writer, nil)
}
