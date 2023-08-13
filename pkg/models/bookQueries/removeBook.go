package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

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
