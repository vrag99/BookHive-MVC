package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"BookHive/pkg/views"
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

	db, _ := models.Connection()
	rows := utils.ExecSql(db, "select * from users where username=?", username)
	defer rows.Close()
	defer db.Close()

	if !rows.Next() {
		loginErr(w, r, types.Err{ErrMsg: "User doesn't exist"})
	} else {
		var usrData types.UserData
		err := rows.Scan(&usrData.Id, &usrData.Username, &usrData.Admin, &usrData.Hash)
		if err != nil {
			panic(err)
		}

		passMatch := utils.MatchPassword(password, usrData.Hash)
		if passMatch {
			token := utils.GenerateJWT(usrData)

			cookie := http.Cookie{
				Name:    "access-token",
				Value:   token,
				Expires: time.Now().Add(48 * time.Hour),
				Path:    "/",
			}

			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/userDashboard", http.StatusSeeOther)
		} else {
			loginErr(w, r, types.Err{ErrMsg: "Incorrect password"})
		}
	}

}

func loginErr(w http.ResponseWriter, r *http.Request, err types.Err) {
	t := views.LoginPage()
	t.Execute(w, err)
}

// func authenticated(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Authenticated, %s!", r.URL.Path[1:])
// }
