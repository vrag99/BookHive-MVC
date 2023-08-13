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
	"strconv"

	"github.com/gorilla/mux"
)

func AdminViews(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
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
		success := bookQueries.RemoveBook(db, id, removeQuantity)
		if success {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	// When admin adds a new book
	booksUpdated := r.FormValue("booksUpdated") == "true"
	invalidBookEntry := r.FormValue("invalidBookEntry") == "true"

	var msg string
	if booksUpdated {
		msg = "booksUpdated"
	} else if invalidBookEntry {
		msg = "invalidBookEntry"
	} else {
		msg = ""
	}
	books := bookQueries.GetAllBooks(db)

	data := types.AdminViewData{
		Username: claims.Username,
		State:    "all",
		Books:    books,
		Error:    msg,
	}

	t := views.AdminDashboard()
	t.Execute(w, data)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	db, _ := models.Connection()
	defer db.Close()

	r.ParseForm()

	bookName := r.FormValue("bookName")
	bookQuantity, _ := strconv.Atoi(r.FormValue("bookQuantity"))

	err := bookQueries.AddBook(db, bookName, bookQuantity)
	if reflect.DeepEqual(err, types.Err{}) {
		http.Redirect(w, r, "/adminDashboard?booksUpdated=true", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/adminDashboard?invalidBookEntry=true", http.StatusSeeOther)
	}

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Printf("invalid request: %s", err)
		fmt.Println(id)
		return
	}

	db, _ := models.Connection()
	defer db.Close()

	success := bookQueries.DeleteBook(db, id)
	if success {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func IssueRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, action := vars["id"], vars["action"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			requestQueries.AcceptIssueRequest(db, id)
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)

		} else if action == "reject" {
			requestQueries.RejectIssueRequest(db, id)
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)
		}
	} else {
		requests := requestQueries.GetIssueRequests(db)

		data := types.UserRequestData{
			Username: claims.Username,
			State:    "issue-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}

func ReturnRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, action := vars["id"], vars["action"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			requestQueries.AcceptReturnRequest(db, id)
			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)

		} else if action == "reject" {
			requestQueries.RejectReturnRequest(db, id)

			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)
		}
	} else {
		requests := requestQueries.GetReturnRequests(db)

		data := types.UserRequestData{
			Username: claims.Username,
			State:    "return-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}

func AdminRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	action, id := vars["action"], vars["id"]

	claims := r.Context().Value(JWTContextKey).(types.Claims)

	db, _ := models.Connection()
	defer db.Close()

	if action != "" && id != "" {
		if action == "accept" {
			requestQueries.AcceptAdminRequest(db, id)
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)
		} else if action == "reject" {
			requestQueries.RejectAdminRequest(db, id)
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)
		}
	} else {
		requests := requestQueries.GetAdminRequests(db)

		data := types.MakeAdminRequestData{
			Username: claims.Username,
			State:    "admin-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}
