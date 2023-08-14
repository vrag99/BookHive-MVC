package requestQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AddReturnRequest(db *sql.DB, bookId string, userId interface{}) error {
	_, err := utils.ExecSql(db, `
		delete from requests
		where status='issued' and bookId=? and userId=?
	`, bookId, userId)
	if err != nil {
		log.Printf("Error deleting upading return request: %v", err)
		return err
	}

	utils.ExecSql(db, `
		insert into requests(status, bookId, userId) 
		values("request-return", ? , ?)
	`, bookId, userId)
	if err != nil {
		log.Printf("Error deleting upading return request: %v", err)
		return err
	}

	return nil
}
