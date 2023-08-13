package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/models/bookQueries"
	"BookHive/pkg/models/requestQueries"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func UserViews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewMode := vars["viewMode"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	data := bookQueries.GetBooksOnViewMode(db, viewMode, claims)
	if reflect.DeepEqual(data, types.UserViewData{}) {
		// If got no data
		fmt.Fprintf(w, "No data. Might be due to an invalid viewMode or error in fetching books.")
	} else {
		t := views.UserDashboard()
		t.Execute(w, data)
	}
}

func RequestBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	requestQueries.AddIssueRequest(db, bookId, claims.Id)

	http.Redirect(w, r, "/userDashboard/requested", http.StatusSeeOther)

}

func RequestReturnBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	requestQueries.AddReturnRequest(db, bookId, claims.Id)

	http.Redirect(w, r, "/userDashboard/toBeReturned", http.StatusSeeOther)
}
