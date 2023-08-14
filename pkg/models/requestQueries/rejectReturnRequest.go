package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func RejectReturnRequest(db *sql.DB, requestId string) error {
	_, err := utils.ExecSql(db, `
		update requests set status = 'issued' 
		where id = ?;
	`, requestId)
	if err != nil {
		log.Printf("Error rejecting return request: %v", err)
		return err
	}

	return nil
}
