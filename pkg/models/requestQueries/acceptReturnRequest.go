package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

func AcceptReturnRequest(db *sql.DB, requestId string) {
	var bookId int
	err := db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		fmt.Printf("Error: '%s' while getting bookId for issue", err)
		panic(err)
	}

	utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)

	utils.ExecSql(db, `
		update books
		set availableQuantity = availableQuantity + 1
		where id = ?;
	`, bookId)
}
