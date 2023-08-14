package requestQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func GetAdminRequests(db *sql.DB) ([]types.MakeAdminRequest, error) {
	rows, err := utils.ExecSql(db, `
		select id, username
		from users
		where requestAdmin = 1
	`)
	if err != nil {
		log.Printf("Error fetching requests for admin access: %v", err)
		return []types.MakeAdminRequest{}, err
	}
	defer rows.Close()

	requests, err := FetchMakeAdminRequests(rows)
	if err != nil {
		return []types.MakeAdminRequest{}, err
	}
	return requests, nil
}
