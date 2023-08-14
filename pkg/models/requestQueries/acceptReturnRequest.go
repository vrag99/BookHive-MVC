package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AcceptReturnRequest(db *sql.DB, requestId string) error {
	var bookId int
	err := db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		log.Printf("Error getting bookId for return: %v", err)
		return err
	}

	_, err = utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)
	if err != nil {
		log.Printf("Error updating requests after accepting return request: %v", err)
		return err
	}

	_, err = utils.ExecSql(db, `
		update books
		set availableQuantity = availableQuantity + 1
		where id = ?;
	`, bookId)
	if err != nil {
		log.Printf("Error updating availableQuantity for books after accepting return request: %v", err)
		return err
	}

	return nil
}
