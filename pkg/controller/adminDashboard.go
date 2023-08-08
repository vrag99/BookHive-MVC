package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"BookHive/pkg/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func AdminViews(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	db, _ := models.Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `select * from books where quantity>=1`)
	defer rows.Close()

	books := models.FetchBooks(rows)

	data := types.AdminViewData{
		Username: claims["username"].(string),
		State:    "all",
		Books:    books,
		Error:    "",
	}

	t := views.AdminDashboard()
	t.Execute(w, data)

}

func AddBook(w http.ResponseWriter, r *http.Request) {

}

func IssueRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	action := vars["action"]

	cookie, _ := r.Cookie("access-token")

	claims , err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	if action != "" && id != "" {
		if action == "accept" {
			models.AcceptIssueRequest(id)
			http.Redirect(w, r, "/adminDashboard/issue-requests", http.StatusSeeOther)

		} else if action == "reject" {
			models.RejectIssueRequest(id)
			http.Redirect(w, r, "/adminDashboard/issue-requests", http.StatusSeeOther)
		}
	} else {
		requests := models.GetIssueRequests()

		data := types.UserRequestData{
			Username: claims["username"].(string),
			State: "issue-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}

func ReturnRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	action := vars["action"]

	cookie, _ := r.Cookie("access-token")

	claims , err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	if action != "" && id != "" {
		if action == "accept" {
			models.AcceptReturnRequest(id)
			http.Redirect(w, r, "/adminDashboard/return-requests", http.StatusSeeOther)
			
			} else if action == "reject" {
				models.RejectReturnRequest(id)
				
			http.Redirect(w, r, "/adminDashboard/return-requests", http.StatusSeeOther)
		}
	} else {
		requests := models.GetReturnRequests()

		data := types.UserRequestData{
			Username: claims["username"].(string),
			State: "return-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}
