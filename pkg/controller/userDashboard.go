package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"BookHive/pkg/views"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func UserViews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewMode := vars["viewMode"]

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	data := models.GetBooksOnViewMode(viewMode, claims)
	if reflect.DeepEqual(data, types.UserViewData{}){
		// If got no data
		fmt.Fprintf(w, "No data. Might be due to an invalid viewMode or error in fetching books.")
	} else{
		t := views.UserDashboard()
		t.Execute(w, data)
	}

}

func RequestBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	db, _ := models.Connection()
	defer db.Close()

	utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-issue", ? , ?)
	`, bookId, claims["id"])

	http.Redirect(w, r, "/userDashboard/requested", http.StatusSeeOther)

}

func RequestReturnBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	db, _ := models.Connection()
	defer db.Close()

	utils.ExecSql(db, `
		delete from requests
		where status='issued' and bookId=? and userId=?
	`, bookId, claims["id"])

	utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-return", ? , ?)
	`, bookId, claims["id"])

	http.Redirect(w, r, "/userDashboard/requested", http.StatusSeeOther)
}
