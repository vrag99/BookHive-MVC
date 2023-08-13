package bookQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func AddBook(db *sql.DB, bookName string, bookQuantity int) types.Err {
	var bookId int
	bookExists := true
	err := db.QueryRow("select id from books where bookName = ?", bookName).Scan(&bookId)
	if err == sql.ErrNoRows {
		bookExists = false
	}

	if bookQuantity > 0 {
		if bookExists {
			utils.ExecSql(db, `
				update books
				set
				quantity = quantity + ?, availableQuantity = availableQuantity + ?
				where id = ?
			`, bookQuantity, bookQuantity, bookId)

			return types.Err{}

		} else {
			utils.ExecSql(db, `
				insert into books (bookName, quantity, availableQuantity) 
				values (?, ?, ?);
			`, bookName, bookQuantity, bookQuantity)

			return types.Err{}
		}
	} else {
		return types.Err{ErrMsg: "invalidQuantity"}
	}
}
