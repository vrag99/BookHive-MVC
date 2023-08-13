package bookQueries

import (
	"BookHive/pkg/types"
	"database/sql"
	"fmt"
)

func FetchBooks(rows *sql.Rows) []types.Book {
	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Name, &book.Quantity, &book.AvailableQuantity)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchBooks = append(fetchBooks, book)
	}
	return fetchBooks
}
