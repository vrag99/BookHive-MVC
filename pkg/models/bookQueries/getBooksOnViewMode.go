package bookQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func GetBooksOnViewMode(db *sql.DB, viewMode string, claims types.Claims) types.UserViewData {
	if viewMode == "requested" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'request-issue' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}

	} else if viewMode == "issued" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'issued' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}

	} else if viewMode == "toBeReturned" {
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			inner join requests r on b.id = r.bookId
			where r.status = 'request-return' and r.userId = ? and b.availableQuantity>=1;
		`, claims.Id)
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims.Username,
			State:    "to-be-returned",
			Books:    books,
		}

	} else {
		viewMode := "available"
		rows := utils.ExecSql(db, `
			select b.*
			from books b
			left join requests r on b.id = r.bookId
			and r.userId = ?
			where r.id is NULL and b.availableQuantity>=0;
		`, claims.Id)
		defer rows.Close()

		books := FetchBooks(rows)
		return types.UserViewData{
			Username: claims.Username,
			State:    viewMode,
			Books:    books,
		}
	}
}
