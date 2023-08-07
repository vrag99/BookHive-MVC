package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"net/http"
	"reflect"
	"time"
)

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	t := views.RegisterPage()
	t.Execute(w, nil)
}

func SignUpRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Getting all the form values
	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")
	registerAsAdmin := r.FormValue("registerAsAdmin")
	adminPassword := r.FormValue("adminPassword")

	var regAsAdmin bool
	if registerAsAdmin == "on" {
		regAsAdmin = true
	} else {
		regAsAdmin = false
	}

	err := models.AddUser(username, password, confirmPassword, regAsAdmin, adminPassword)
	if !reflect.DeepEqual(err, types.Err{}) {
		signUpErr(w, r, err)
	} else {
		cookie := http.Cookie{
			Name: "access-token",
			Value: "",
			Expires: time.Now(),
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/login?registered=true", http.StatusSeeOther)
	}

}

func signUpErr(w http.ResponseWriter, r *http.Request, err types.Err) {
	t := views.RegisterPage()
	t.Execute(w, err)
}
