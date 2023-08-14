package bookQueries

import (
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func AppendBook(db *sql.DB, bookId int, quantity int) error {
	_, err := utils.ExecSql(db, `
		update books 
		set quantity = quantity + ?, availableQuantity = availableQuantity + ?
		where id = ?
	`, quantity, quantity, bookId)
	if err != nil {
		log.Printf("Error appending book: %v", err)
		return err
	}

	return nil
}
