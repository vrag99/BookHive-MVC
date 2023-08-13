package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	t := views.NotFound()
	t.Execute(w, nil)
}
