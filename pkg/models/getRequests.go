package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
	"fmt"
)

func FetchRequests(rows *sql.Rows) []types.UserRequest {
	var fetchRequests []types.UserRequest
	for rows.Next() {
		var request types.UserRequest
		err := rows.Scan(&request.Id, &request.Username, &request.BookName)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests
}

func GetIssueRequests() []types.UserRequest {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `
		select requests.id , users.username, books.book_name
		from requests
		join books on requests.book_id = books.id
		join users on requests.user_id = users.id
		where requests.status = 'request-issue' and available_qty>=1;
	`)
	defer rows.Close()

	requests := FetchRequests(rows)
	return requests
}

func GetReturnRequests() []types.UserRequest {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db , `
		select requests.id , users.username, books.book_name
		from requests
		join books on requests.book_id = books.id
		join users on requests.user_id = users.id
		where requests.status = 'request-return';
	`)

	defer rows.Close()
	requests := FetchRequests(rows)

	return requests
}

func AcceptIssueRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
	update requests set status = 'issued' 
	where id = ?;
	`, requestId)

	var bookId int
	err := db.QueryRow("select r.book_id from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		fmt.Printf("Error: '%s' while getting bookId for issue", err)
	}

	utils.ExecSql(db, `
		update books
		set available_qty = available_qty - 1
		where id = ?;
	`, bookId)
}

func AcceptReturnRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	var bookId int
	err := db.QueryRow("select r.book_id from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		fmt.Printf("Error: '%s' while getting bookId for issue", err)
		panic(err)
	}

	utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)

	utils.ExecSql(db, `
		update books
		set available_qty = available_qty + 1
		where id = ?;
	`, bookId)
}

func RejectIssueRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
		delete from requests
		where id = ?;
	`, requestId)
}

func RejectReturnRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
		update requests set status = 'issued' 
		where id = ?;
	`, requestId)
}
