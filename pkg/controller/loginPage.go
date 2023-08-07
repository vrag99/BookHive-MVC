package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"BookHive/pkg/views"
	"reflect"
	"strconv"

	"fmt"
	"net/http"
	"time"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	registered, _ := strconv.ParseBool(r.FormValue("registered"))

	cookie, err := r.Cookie("access-token")
	if err != nil {
		data := struct {
			ErrMsg     string
			Registered bool
		}{
			ErrMsg:     "",
			Registered: registered,
		}
		t := views.LoginPage()
		t.Execute(w, data)
	} else {
		claims, err := utils.DecodeJWT(cookie.Value)
		if err != nil {
			fmt.Println("Invalid JWT")
			panic(err)
		} else {
			isAdmin, _ := strconv.ParseBool(fmt.Sprintf("%d", claims["admin"]))
			if isAdmin {
				http.Redirect(w, r, "/adminDashboard", http.StatusSeeOther)
			} else {
				http.Redirect(w, r, "/userDashboard", http.StatusSeeOther)
			}
		}
	}
}

func LoginRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	token, err, isAdmin := models.GetJWT(username, password)
	if !reflect.DeepEqual(err, types.Err{}) {
		loginErr(w, r, err)
	} else {

		cookie := http.Cookie{
			Name:    "access-token",
			Value:   token,
			Expires: time.Now().Add(48 * time.Hour),
			Path:    "/",
		}

		http.SetCookie(w, &cookie)
		if isAdmin {
			http.Redirect(w, r, "/adminDashboard", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/userDashboard", http.StatusSeeOther)
		}
	}

}

func loginErr(w http.ResponseWriter, r *http.Request, err types.Err) {
	t := views.LoginPage()
	t.Execute(w, err)
}
