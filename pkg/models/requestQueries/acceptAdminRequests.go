package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AcceptAdminRequest(db *sql.DB, userId string) error {
	_, err := utils.ExecSql(db, `
		update users
		set admin = 1, requestAdmin = 0
		where id = ?
	`, userId)

	if err != nil {
		log.Printf("Error occured while accepting admin request: %v", err)
		return err
	}

	return nil
}
