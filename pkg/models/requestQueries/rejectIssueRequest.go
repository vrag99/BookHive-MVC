package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func RejectIssueRequest(db *sql.DB, requestId string) {
	utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)
}
