package bookQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func GetBooksOnViewMode(db *sql.DB, viewMode string, claims types.Claims) (types.UserViewData, error) {
	if viewMode == "requested" {
		rows, err := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'request-issue' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		if err != nil {
			log.Printf("Error getting the requested books for issue: %v", err)
			return types.UserViewData{}, err
		}
		defer rows.Close()

		books, err := FetchBooks(rows)
		if err != nil {
			return types.UserViewData{}, err
		}
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}, nil

	} else if viewMode == "issued" {
		rows, err := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'issued' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		if err != nil {
			log.Printf("Error getting the issued books: %v", err)
			return types.UserViewData{}, err
		}
		defer rows.Close()

		books, err := FetchBooks(rows)
		if err != nil {
			return types.UserViewData{}, err
		}
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}, nil

	} else if viewMode == "toBeReturned" {
		rows, err := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'request-return' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		if err != nil {
			log.Printf("Error getting the requested books for return: %v", err)
			return types.UserViewData{}, err
		}
		defer rows.Close()

		books, err := FetchBooks(rows)
		if err != nil {
			return types.UserViewData{}, err
		}
		return types.UserViewData{
			Username: claims.Username,
			State:    "to-be-returned",
			Books:    books,
		}, nil

	} else {
		viewMode := "available"
		rows, err := utils.ExecSql(db, `
			select b.*
			from books b
			left join requests r on b.id = r.bookId
			and r.userId = ?
			where r.id is NULL and b.availableQuantity>=0;
		`, claims.Id)
		if err != nil {
			log.Printf("Error getting the available books: %v", err)
			return types.UserViewData{}, err
		}
		defer rows.Close()

		books, err := FetchBooks(rows)
		if err != nil {
			return types.UserViewData{}, err
		}
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}, nil
	}
}
