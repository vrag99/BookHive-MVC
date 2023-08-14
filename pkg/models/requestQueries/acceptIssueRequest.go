package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AcceptIssueRequest(db *sql.DB, requestId string) error {
	_, err := utils.ExecSql(db, `
		update requests set status = 'issued' 
		where id = ?;
	`, requestId)
	if err != nil {
		log.Printf("Error updating issue request: %v", err)
		return err
	}

	var bookId int
	err = db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		log.Printf("Error getting bookId for issue: %v", err)
		return err
	}

	_, err = utils.ExecSql(db, `
		update books
		set availableQuantity = availableQuantity - 1
		where id = ?;
	`, bookId)
	if err != nil {
		log.Printf("Error updating availableQuantity after issue: %v", err)
		return err
	}

	return nil
}
