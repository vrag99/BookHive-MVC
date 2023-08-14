package requestQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"log"
)

func FetchMakeAdminRequests(rows *sql.Rows) ([]types.MakeAdminRequest, error) {
	var fetchRequests []types.MakeAdminRequest
	for rows.Next() {
		var request types.MakeAdminRequest
		err := rows.Scan(&request.Id, &request.Username)
		if err != nil {
			log.Printf("Error scanning requests for admin access: %v", err)
			return []types.MakeAdminRequest{}, err
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests, nil
}
