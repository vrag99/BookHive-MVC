package requestQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"fmt"
)

func FetchRequests(rows *sql.Rows) []types.UserRequest {
	var fetchRequests []types.UserRequest
	for rows.Next() {
		var request types.UserRequest
		err := rows.Scan(&request.Id, &request.Username, &request.BookName)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests
}
