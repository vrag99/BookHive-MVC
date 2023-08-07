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

func UserViews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	viewMode := vars["viewMode"]

	cookie, _ := r.Cookie("access-token")

	claims, err := utils.DecodeJWT(cookie.Value)
	if err != nil {
		fmt.Println("Invalid JWT token")
		return
	}

	db, _ := models.Connection()
	defer db.Close()

	// w.Write([]byte("ye lo username: " + claims["username"].(string)))

	if viewMode == "requested" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'request-issue' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		var fetchBooks []types.Book
		for rows.Next() {
			var book types.Book
			err := rows.Scan(&book.Id, &book.Name, &book.Qty, &book.AvailableQty)
			if err != nil {
				fmt.Println("Error fetching books for userViews")
				panic(err)
			}
			fetchBooks = append(fetchBooks, book)
		}

		data := types.UserViewData{
			State: viewMode,
			Books: fetchBooks,
		}
		t := views.UserDashboard()
		t.Execute(w, data)

	} else if viewMode == "issued" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'issued' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		var fetchBooks []types.Book
		for rows.Next() {
			var book types.Book
			err := rows.Scan(&book.Id, &book.Name, &book.Qty, &book.AvailableQty)
			if err != nil {
				fmt.Println("Error fetching books for userViews")
				panic(err)
			}
			fetchBooks = append(fetchBooks, book)
		}

		data := types.UserViewData{
			State: viewMode,
			Books: fetchBooks,
		}
		t := views.UserDashboard()
		t.Execute(w, data)

	} else if viewMode == "to-be-returned" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'issued' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		var fetchBooks []types.Book
		for rows.Next() {
			var book types.Book
			err := rows.Scan(&book.Id, &book.Name, &book.Qty, &book.AvailableQty)
			if err != nil {
				fmt.Println("Error fetching books for userViews")
				panic(err)
			}
			fetchBooks = append(fetchBooks, book)
		}

		data := types.UserViewData{
			State: viewMode,
			Books: fetchBooks,
		}
		t := views.UserDashboard()
		t.Execute(w, data)

	} else {
		viewMode := "available"
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			left join requests r on b.id = r.book_id
			and r.user_id = ?
			where r.id is NULL and available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		var fetchBooks []types.Book
		for rows.Next() {
			var book types.Book
			err := rows.Scan(&book.Id, &book.Name, &book.Qty, &book.AvailableQty)
			if err != nil {
				fmt.Println("Error fetching books for userViews")
				panic(err)
			}
			fetchBooks = append(fetchBooks, book)
		}

		data := types.UserViewData{
			Username: claims["username"].(string),
			State:    viewMode,
			Books:    fetchBooks,
		}

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
		insert into requests(status, book_id, user_id) 
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
		insert into requests(status, book_id, user_id) 
		values("request-return", ? , ?)
	`, bookId, claims["id"])

	http.Redirect(w, r, "/userDashboard/requested", http.StatusSeeOther)
}
