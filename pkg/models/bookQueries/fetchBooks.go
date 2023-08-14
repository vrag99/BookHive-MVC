package bookQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"log"
)

func FetchBooks(rows *sql.Rows) ([]types.Book, error) {
	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Quantity, &book.AvailableQuantity)
		if err != nil {
			log.Printf("Error fetching books: %v", err)
			return []types.Book{}, nil
		}
		fetchBooks = append(fetchBooks, book)
	}
	return fetchBooks, nil
}
