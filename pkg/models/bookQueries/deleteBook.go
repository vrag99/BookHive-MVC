package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

func DeleteBook(db *sql.DB, bookId int) bool {
	var requests int
	err := db.QueryRow(`select count(*) from requests r where r.bookId = ?`, bookId).Scan(&requests)
	if err != nil {
		fmt.Printf("Error: '%s' while getting requests for the book.", err)
		return false
	}

	if requests == 0 {
		// The book isn't requested, safe to delete.
		utils.ExecSql(db, `
			delete from books
			where id = ?;
		`, bookId)
		return true
	} else {
		return false
	}
}
