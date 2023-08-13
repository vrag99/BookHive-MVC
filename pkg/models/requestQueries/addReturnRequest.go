package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func AddReturnRequest(db *sql.DB, bookId string, userId interface{}) {
	utils.ExecSql(db, `
		delete from requests
		where status='issued' and bookId=? and userId=?
	`, bookId, userId)

	utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-return", ? , ?)
	`, bookId, userId)
}
