package controller

import (
	"BookHive/pkg/middleware"
	"BookHive/pkg/models"
	"BookHive/pkg/models/bookQueries"
	"BookHive/pkg/types"
	"BookHive/pkg/views"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func AdminViews(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.JWTContextKey).(types.Claims)

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	defer db.Close()

	// When the quantity of an existing book is changed.
	// Passing params through axios.
	id, _ := strconv.Atoi(r.FormValue("id"))
	addedQuantity, _ := strconv.Atoi(r.FormValue("addedQuantity"))
	removeQuantity, _ := strconv.Atoi(r.FormValue("removeQuantity"))

	if addedQuantity > 0 {
		bookQueries.AppendBook(db, id, addedQuantity)
		w.WriteHeader(http.StatusOK)

	} else if removeQuantity > 0 {
		ok, err := bookQueries.RemoveBook(db, id, removeQuantity)
		if err != nil {
			http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
		}

		if ok {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	// When admin adds a new book
	booksUpdated := r.FormValue("booksUpdated") == "true"
	invalidBookEntry := r.FormValue("invalidBookEntry") == "true"

	var message string
	if booksUpdated {
		message = "booksUpdated"
	} else if invalidBookEntry {
		message = "invalidBookEntry"
	} else {
		message = ""
	}

	books, err := bookQueries.GetAllBooks(db)
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	data := types.AdminViewData{
		Username: claims.Username,
		State:    "all",
		Books:    books,
		Error:    message,
	}

	t := views.Mode("adminDashboard")
	t.Execute(w, data)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	r.ParseForm()

	bookName := r.FormValue("bookName")
	bookQuantity, _ := strconv.Atoi(r.FormValue("bookQuantity"))

	error := bookQueries.AddBook(db, bookName, bookQuantity)
	if reflect.DeepEqual(error, types.Err{}) {
		http.Redirect(w, r, "/adminDashboard?booksUpdated=true", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/adminDashboard?invalidBookEntry=true", http.StatusSeeOther)
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("An invalid number was passed as id for book: %v", err)
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	db, err := models.Connection()
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}
	defer db.Close()

	ok, err := bookQueries.DeleteBook(db, id)
	if err != nil {
		http.Redirect(w, r, "/internalServerError", http.StatusSeeOther)
	}

	if ok {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
