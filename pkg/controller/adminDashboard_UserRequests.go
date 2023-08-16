package controller

import (
	"BookHive/pkg/middleware"
	"BookHive/pkg/models"
	"BookHive/pkg/models/requestQueries"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"net/http"

	"github.com/gorilla/mux"
)

func IssueRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, action := vars["id"], vars["action"]

	claims := r.Context().Value(middleware.JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			err = requestQueries.AcceptIssueRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)

		} else if action == "reject" {
			err = requestQueries.RejectIssueRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)
		}

	} else {
		requests, err := requestQueries.GetIssueRequests(db)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
		}

		data := types.UserRequestData{
			Username: claims.Username,
			State:    "issue-requests",
			Requests: requests,
		}

		t := views.Mode("adminDashboard")
		t.Execute(w, data)
	}
}

func ReturnRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, action := vars["id"], vars["action"]

	claims := r.Context().Value(middleware.JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			err = requestQueries.AcceptReturnRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)

		} else if action == "reject" {
			err = requestQueries.RejectReturnRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)
		}
	} else {
		requests, err := requestQueries.GetReturnRequests(db)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
		}

		data := types.UserRequestData{
			Username: claims.Username,
			State:    "return-requests",
			Requests: requests,
		}

		t := views.Mode("adminDashboard")
		t.Execute(w, data)
	}
}

func AdminRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action, id := vars["action"], vars["id"]

	claims := r.Context().Value(middleware.JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			err = requestQueries.AcceptAdminRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)

		} else if action == "reject" {
			err = requestQueries.RejectAdminRequest(db, id)
			if err != nil {
				http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
			}
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)
		}

	} else {
		requests, err := requestQueries.GetAdminRequests(db)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
		}

		data := types.MakeAdminRequestData{
			Username: claims.Username,
			State:    "admin-requests",
			Requests: requests,
		}

		t := views.Mode("adminDashboard")
		t.Execute(w, data)
	}
}
