package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func FetchBooks(rows *sql.Rows) []types.Book {
	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Qty, &book.AvailableQty)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchBooks = append(fetchBooks, book)
	}
	return fetchBooks
}

func GetBooksOnViewMode(viewMode string, claims jwt.MapClaims) types.UserViewData {
	db, _ := Connection()
	defer db.Close()

	if viewMode == "requested" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'request-issue' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: viewMode,
			Books: books,
		}

	} else if viewMode == "issued" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'issued' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: viewMode,
			Books: books,
		}

	} else if viewMode == "to-be-returned" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.book_id
			where r.status = 'issued' and r.user_id = ? and b.available_qty>=1;
		`, claims["id"])
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: viewMode,
			Books: books,
		}

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

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: viewMode,
			Books: books,
		}
	}
}

func GetAllBooks() []types.Book {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `select * from books where quantity>=1`)
	defer rows.Close()

	books := FetchBooks(rows)
	return books
}