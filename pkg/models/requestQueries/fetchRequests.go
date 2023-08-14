package requestQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"log"
)

func FetchRequests(rows *sql.Rows) ([]types.UserRequest, error) {
	var fetchRequests []types.UserRequest
	for rows.Next() {
		var request types.UserRequest
		err := rows.Scan(&request.Id, &request.Username, &request.BookName)
		if err != nil {
			log.Printf("Error scanning requests for books: %v", err)
			return []types.UserRequest{}, err
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests, nil
}
