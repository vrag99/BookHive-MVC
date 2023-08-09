package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

func AddBook(bookName string, bookQty int) types.Err {
	db, _ := Connection()
	defer db.Close()

	var bookId int
	bookExists := true
	err := db.QueryRow("select id from books where book_name = ?", bookName).Scan(&bookId)
	if err == sql.ErrNoRows {
		bookExists = false
	}

	if bookQty > 0 {
		if bookExists {
			utils.ExecSql(db, `
				update books
				set
				quantity = quantity + ?, available_qty = available_qty + ?
				where id = ?
			`, bookQty, bookQty, bookId)

			return types.Err{}

		} else {
			utils.ExecSql(db, `
				insert into books (book_name, quantity, available_qty) 
				values (?, ?, ?);
			`, bookName, bookQty, bookQty)

			return types.Err{}
		}
	} else {
		return types.Err{ErrMsg: "invalidQty"}
	}
}

func AppendBook(bookId int, quantity int) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
		update books 
		set quantity = quantity + ?, available_qty = available_qty + ?
		where id = ?
	`, quantity, quantity, bookId)
}

func RemoveBook(bookId int, quantity int) bool {
	db, _ := Connection()
	defer db.Close()

	var availableQty, pendingRequests int
	err := db.QueryRow(`select available_qty from books where books.id = ?`, bookId).Scan(&availableQty)
	if err != nil {
		fmt.Printf("Error: '%s' while getting available quantity for book", err)
		panic(err)
	}

	err = db.QueryRow(`select count(*) from requests r where r.book_id = ? and r.status != 'issued'`, bookId).Scan(&pendingRequests)
	if err != nil {
		fmt.Printf("Error: '%s' while getting pending requests for the book.", err)
		panic(err)
	}

	if availableQty - pendingRequests >= quantity {
		utils.ExecSql(db, `
			update books 
			set quantity = quantity - ?, available_qty = available_qty - ?
			where id = ?;
		`, quantity, quantity, bookId)
		return true
	} else {
		return false
	}

}
