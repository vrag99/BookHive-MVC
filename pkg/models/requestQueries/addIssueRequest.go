package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AddIssueRequest(db *sql.DB, bookId string, userId interface{}) (error) {
	_, err := utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-issue", ? , ?)
	`, bookId, userId)
	if err != nil {
		log.Printf("Error adding issue request: %v", err)
		return err
	}

	return nil
}
