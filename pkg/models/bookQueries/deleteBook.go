package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func DeleteBook(db *sql.DB, bookId int) (bool, error) {
	var requests int
	err := db.QueryRow(`select count(*) from requests r where r.bookId = ?`, bookId).Scan(&requests)
	if err != nil {
		log.Printf("Error: '%s' while getting requests for the book.", err)
		return false, err
	}

	if requests == 0 {
		// The book isn't requested, safe to delete.
		_, err := utils.ExecSql(db, `
			delete from books
			where id = ?;
		`, bookId)
		if err!= nil {
            log.Printf("Error: '%s' while deleting the book.", err)
            return false, err
        }
		return true, nil

	} else {

		return false, nil
	}
}
