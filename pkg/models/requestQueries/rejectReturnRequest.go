package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
)

func RejectReturnRequest(db *sql.DB, requestId string) {
	utils.ExecSql(db, `
		update requests set status = 'issued' 
		where id = ?;
	`, requestId)
}
