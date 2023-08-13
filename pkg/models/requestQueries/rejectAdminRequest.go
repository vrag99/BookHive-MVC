package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func RejectAdminRequest(db *sql.DB, userId string) {
	utils.ExecSql(db, `
		update users
		set admin = 0, requestAdmin = 0
		where id = ?
	`, userId)
}
