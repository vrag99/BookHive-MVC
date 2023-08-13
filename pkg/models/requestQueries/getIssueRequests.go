package requestQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func GetIssueRequests(db *sql.DB) []types.UserRequest {
	rows := utils.ExecSql(db, `
		select requests.id , users.username, books.bookName
		from requests
		join books on requests.bookId = books.id
		join users on requests.userId = users.id
		where requests.status = 'request-issue' and availableQuantity>=1;
	`)
	defer rows.Close()

	requests := FetchRequests(rows)
	return requests
}
