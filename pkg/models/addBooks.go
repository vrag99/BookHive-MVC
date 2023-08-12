package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
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

func AppendBook(db *sql.DB, bookId int, quantity int) {
	utils.ExecSql(db, `
		update books 
		set quantity = quantity + ?, availableQuantity = availableQuantity + ?
		where id = ?
	`, quantity, quantity, bookId)
}

func RemoveBook(db *sql.DB, bookId int, quantity int) bool {
	var availableQuantity, pendingRequests int

	err := db.QueryRow(`select availableQuantity from books where books.id = ?`, bookId).Scan(&availableQuantity)
	if err != nil {
		fmt.Printf("Error: '%s' while getting available quantity for book", err)
		panic(err)
	}

	err = db.QueryRow(`select count(*) from requests r where r.bookId = ? and r.status != 'issued'`, bookId).Scan(&pendingRequests)
	if err != nil {
		fmt.Printf("Error: '%s' while getting pending requests for the book.", err)
		panic(err)
	}

	if availableQuantity-pendingRequests >= quantity {
		utils.ExecSql(db, `
			update books 
			set quantity = quantity - ?, availableQuantity = availableQuantity - ?
			where id = ?;
		`, quantity, quantity, bookId)
		return true
	} else {
		return false
	}

}
