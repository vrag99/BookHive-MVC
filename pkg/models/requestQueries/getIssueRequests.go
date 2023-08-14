package requestQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func GetIssueRequests(db *sql.DB) ([]types.UserRequest, error) {
	rows, err := utils.ExecSql(db, `
		select requests.id , users.username, books.bookName
		from requests
		join books on requests.bookId = books.id
		join users on requests.userId = users.id
		where requests.status = 'request-issue' and availableQuantity>=1;
	`)
	if err != nil {
		log.Printf("Error fetching issue requests: %v", err)
		return []types.UserRequest{}, err
	}
	defer rows.Close()

	requests, err := FetchRequests(rows)
	if err != nil {
		return []types.UserRequest{}, err
	}

	return requests, nil
}
