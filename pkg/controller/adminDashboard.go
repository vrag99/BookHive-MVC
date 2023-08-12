package controller

import (
	"BookHive/pkg/models"
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"BookHive/pkg/views"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func AdminViews(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	// When the quantity of an existing book is changed.
	// Passing params through axios.
	id, _ := strconv.Atoi(r.FormValue("id"))
	addedQuantity, _ := strconv.Atoi(r.FormValue("addedQuantity"))
	removeQuantity, _ := strconv.Atoi(r.FormValue("removeQuantity"))

	if addedQuantity > 0 {
		models.AppendBook(id, addedQuantity)
		w.WriteHeader(http.StatusOK)

	} else if removeQuantity > 0 {
		success := models.RemoveBook(id, removeQuantity)
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
	books := models.GetAllBooks()

	data := types.AdminViewData{
		Username: claims["username"].(string),
		State:    "all",
		Books:    books,
		Error:    msg,
	}

	t := views.AdminDashboard()
	t.Execute(w, data)

}

func AddBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	bookName := r.FormValue("bookName")
	bookQuantity, _ := strconv.Atoi(r.FormValue("bookQuantity"))

	err := models.AddBook(bookName, bookQuantity)
	if reflect.DeepEqual(err, types.Err{}) {
		http.Redirect(w, r, "/adminDashboard?booksUpdated=true", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/adminDashboard?invalidBookEntry=true", http.StatusSeeOther)
	}

}

func IssueRequests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, action := vars["id"], vars["action"]

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	if action != "" && id != "" {
		if action == "accept" {
			models.AcceptIssueRequest(id)
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)

		} else if action == "reject" {
			models.RejectIssueRequest(id)
			http.Redirect(w, r, "/adminDashboard/issueRequests", http.StatusSeeOther)
		}
	} else {
		requests := models.GetIssueRequests()

		data := types.UserRequestData{
			Username: claims["username"].(string),
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

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	if action != "" && id != "" {
		if action == "accept" {
			models.AcceptReturnRequest(id)
			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)

		} else if action == "reject" {
			models.RejectReturnRequest(id)

			http.Redirect(w, r, "/adminDashboard/returnRequests", http.StatusSeeOther)
		}
	} else {
		requests := models.GetReturnRequests()

		data := types.UserRequestData{
			Username: claims["username"].(string),
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

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	if action != "" && id != "" {
		if action == "accept" {
			models.AcceptAdminRequest(id)
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)
		} else if action == "reject" {
			models.RejectAdminRequest(id)
			http.Redirect(w, r, "/adminDashboard/adminRequests", http.StatusSeeOther)
		}
	} else {
		requests := models.GetAdminRequests()

		data := types.MakeAdminRequestData{
			Username: claims["username"].(string),
			State:    "admin-requests",
			Requests: requests,
		}

		t := views.AdminDashboard()
		t.Execute(w, data)
	}
}
