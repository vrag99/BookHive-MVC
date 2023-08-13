package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func AppendBook(db *sql.DB, bookId int, quantity int) {
	utils.ExecSql(db, `
		update books 
		set quantity = quantity + ?, availableQuantity = availableQuantity + ?
		where id = ?
	`, quantity, quantity, bookId)
}
