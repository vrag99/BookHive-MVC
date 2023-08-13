package bookQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func GetAllBooks(db *sql.DB) []types.Book {
	rows := utils.ExecSql(db, `select * from books`)
	defer rows.Close()

	books := FetchBooks(rows)
	return books
}
