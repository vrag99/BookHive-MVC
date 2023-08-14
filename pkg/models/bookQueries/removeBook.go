package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
	"log"
)

func RemoveBook(db *sql.DB, bookId int, quantity int) (bool, error) {
	var availableQuantity, pendingRequests int

	err := db.QueryRow(`select availableQuantity from books where books.id = ?`, bookId).Scan(&availableQuantity)
	if err != nil {
		log.Printf("Error getting available quantity for book: %v", err)
		return false, err
	}

	err = db.QueryRow(`select count(*) from requests r where r.bookId = ? and r.status != 'issued'`, bookId).Scan(&pendingRequests)
	if err != nil {
		fmt.Printf("Error getting pending requests for the book: %s", err)
		return false, err
	}

	if availableQuantity-pendingRequests >= quantity {
		_, err := utils.ExecSql(db, `
			update books 
			set quantity = quantity - ?, availableQuantity = availableQuantity - ?
			where id = ?;
		`, quantity, quantity, bookId)
		if err != nil {
			log.Printf("Error updating the quantity for book: %v", err)
			return false, err
		}

		return true, nil
	} else {

		return false, nil
	}

}
