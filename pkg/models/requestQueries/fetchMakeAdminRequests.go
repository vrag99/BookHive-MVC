package requestQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"fmt"
)

func FetchMakeAdminRequests(rows *sql.Rows) []types.MakeAdminRequest {
	var fetchRequests []types.MakeAdminRequest
	for rows.Next() {
		var request types.MakeAdminRequest
		err := rows.Scan(&request.Id, &request.Username)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests

}
