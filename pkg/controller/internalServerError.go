package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	t := views.InternalServerError()
	t.Execute(w, nil)
}
