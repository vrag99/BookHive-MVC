package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func AcceptAdminRequest(db *sql.DB, userId string) {
	utils.ExecSql(db, `
		update users
		set admin = 1, requestAdmin = 0
		where id = ?
	`, userId)
}
