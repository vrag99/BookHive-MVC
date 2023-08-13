package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func ForbiddenRequest(w http.ResponseWriter, r *http.Request) {
	t := views.ForbiddenRequest()
	t.Execute(w, nil)
}
