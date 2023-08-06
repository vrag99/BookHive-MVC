package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"BookHive/pkg/utils"
	"net/http"
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
	
	var regAsAdmin bool
	if registerAsAdmin == "on" {regAsAdmin = true} else {regAsAdmin = false}
	adminPassword := r.FormValue("adminPassword")

	config, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
	ADMIN_PASS := config.AdminPassword

	var adminApproved bool
	if regAsAdmin {
		if adminPassword == ADMIN_PASS {adminApproved = true} else {adminApproved = false}
	} else {
		adminApproved = false
	}

	db, _ := models.Connection()
	rows := utils.ExecSql(db, "select * from users where username=?", username)
	rows.Close()
	defer db.Close()

	if rows.Next() {
		signUpErr(w, r, types.Err{ErrMsg: "User already exists"})
	} else if password != confirmPassword {
		signUpErr(w, r, types.Err{ErrMsg: "The passwords don't match"})
	} else if regAsAdmin && !adminApproved {
		signUpErr(w, r, types.Err{ErrMsg: "Incorrect admin passcode"})
	} else {
		hashedPassword := utils.HashPassword(password)
		utils.ExecSql(db, "insert into users (username, admin, hash) values(?, ?, ?)", username, regAsAdmin, hashedPassword)

		cookie := http.Cookie{
			Name:    "access-token",
			Value:   "",
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
