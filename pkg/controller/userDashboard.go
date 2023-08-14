package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/models/bookQueries"
	"BookHive/pkg/models/requestQueries"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"net/http"

	"github.com/gorilla/mux"
)

func UserViews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewMode := vars["viewMode"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	data, err := bookQueries.GetBooksOnViewMode(db, viewMode, claims)
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	t := views.UserDashboard()
	t.Execute(w, data)
}

func RequestBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	err = requestQueries.AddIssueRequest(db, bookId, claims.Id)
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/userDashboard/requested", http.StatusSeeOther)

}

func RequestReturnBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	err = requestQueries.AddReturnRequest(db, bookId, claims.Id)
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/userDashboard/toBeReturned", http.StatusSeeOther)
}
