package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func RejectIssueRequest(db *sql.DB, requestId string) error {
	_, err := utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)
	if err != nil {
		log.Printf("Error rejecting issue request: %v", err)
		return err
	}

	return nil
}
