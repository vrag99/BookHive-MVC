package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func RejectAdminRequest(db *sql.DB, userId string) error {
	_, err := utils.ExecSql(db, `
		update users
		set admin = 0, requestAdmin = 0
		where id = ?
	`, userId)
	if err != nil {
		log.Printf("Error rejecting request for admin access: %v", err)
		return err
	}

	return nil
}
