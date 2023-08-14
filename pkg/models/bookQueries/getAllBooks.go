package bookQueries

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"log"
)

func GetAllBooks(db *sql.DB) ([]types.Book, error) {
	rows, err := utils.ExecSql(db, `select * from books`)
	if err != nil {
		log.Printf("Error getting the books: %v", err)
		return []types.Book{}, err
	}
	defer rows.Close()

	books, err := FetchBooks(rows)
	if err != nil {
		return []types.Book{}, err
	}
	return books, nil
}
