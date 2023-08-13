package requestQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func GetAdminRequests(db *sql.DB) []types.MakeAdminRequest {
	rows := utils.ExecSql(db, `
		select id, username
		from users
		where requestAdmin = 1
	`)

	requests := FetchMakeAdminRequests(rows)
	return requests
}
