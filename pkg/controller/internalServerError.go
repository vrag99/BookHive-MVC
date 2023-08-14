package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	t := views.Mode("internalServerError")
	t.Execute(w, nil)
}
