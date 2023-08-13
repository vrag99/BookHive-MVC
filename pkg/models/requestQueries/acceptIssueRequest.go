package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

func AcceptIssueRequest(db *sql.DB, requestId string) {
	utils.ExecSql(db, `
	update requests set status = 'issued' 
	where id = ?;
	`, requestId)

	var bookId int
	err := db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		fmt.Printf("Error: '%s' while getting bookId for issue", err)
	}

	utils.ExecSql(db, `
		update books
		set availableQuantity = availableQuantity - 1
		where id = ?;
	`, bookId)
}
