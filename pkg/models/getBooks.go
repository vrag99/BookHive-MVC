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
		err := rows.Scan(&book.Id, &book.Name, &book.Quantity, &book.AvailableQuantity)
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
			inner join requests r on b.id = r.bookId
			where r.status = 'request-issue' and r.userId = ? and b.availableQuantity>=1;
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
			inner join requests r on b.id = r.bookId
			where r.status = 'issued' and r.userId = ? and b.availableQuantity>=1;
		`, claims["id"])
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: viewMode,
			Books: books,
		}

	} else if viewMode == "toBeReturned" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'request-return' and r.userId = ? and b.availableQuantity>=1;
		`, claims["id"])
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims["username"].(string),
			State: "to-be-returned",
			Books: books,
		}

	} else {
		viewMode := "available"
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			left join requests r on b.id = r.bookId
			and r.userId = ?
			where r.id is NULL and b.availableQuantity>=1;
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