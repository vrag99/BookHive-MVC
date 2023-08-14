package controller

import (
	"BookHive/pkg/views"
	"net/http"
)

func ForbiddenRequest(w http.ResponseWriter, r *http.Request) {
	t := views.Mode("forbiddenRequest")
	t.Execute(w, nil)
}
