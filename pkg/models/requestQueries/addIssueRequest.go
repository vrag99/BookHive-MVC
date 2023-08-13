package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func AddIssueRequest(db *sql.DB, bookId string, userId interface{}) {
	utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-issue", ? , ?)
	`, bookId, userId)
}
