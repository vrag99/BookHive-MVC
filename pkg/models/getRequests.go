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

func FetchMakeAdminRequests(rows *sql.Rows) []types.MakeAdminRequest {
	var fetchRequests []types.MakeAdminRequest
	for rows.Next() {
		var request types.MakeAdminRequest
		err := rows.Scan(&request.Id, &request.Username)
		if err != nil {
			fmt.Println("Error fetching books")
			panic(err)
		}
		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests

}

func GetAdminRequests() []types.MakeAdminRequest {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `
		select id, username
		from users
		where requestAdmin = 1
	`)

	requests := FetchMakeAdminRequests(rows)
	return requests
}

func GetIssueRequests() []types.UserRequest {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `
		select requests.id , users.username, books.bookName
		from requests
		join books on requests.bookId = books.id
		join users on requests.userId = users.id
		where requests.status = 'request-issue' and availableQty>=1;
	`)
	defer rows.Close()

	requests := FetchRequests(rows)
	return requests
}

func GetReturnRequests() []types.UserRequest {
	db, _ := Connection()
	defer db.Close()

	rows := utils.ExecSql(db, `
		select requests.id , users.username, books.bookName
		from requests
		join books on requests.bookId = books.id
		join users on requests.userId = users.id
		where requests.status = 'request-return';
	`)

	defer rows.Close()
	requests := FetchRequests(rows)

	return requests
}

func AcceptAdminRequest(userId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
		update users
		set admin = 1, requestAdmin = 0
		where id = ?
	`, userId)
}

func AcceptIssueRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
	update requests set status = 'issued' 
	where id = ?;
	`, requestId)

	var bookId int
	err := db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
	if err != nil {
		fmt.Printf("Error: '%s' while getting bookId for issue", err)
	}

	utils.ExecSql(db, `
		update books
		set availableQty = availableQty - 1
		where id = ?;
	`, bookId)
}

func AcceptReturnRequest(requestId string) {
	db, _ := Connection()
	defer db.Close()

	var bookId int
	err := db.QueryRow("select r.bookId from requests r where r.id = ?", requestId).Scan(&bookId)
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
		set availableQty = availableQty + 1
		where id = ?;
	`, bookId)
}

func RejectAdminRequest(userId string) {
	db, _ := Connection()
	defer db.Close()

	utils.ExecSql(db, `
		update users
		set admin = 0, requestAdmin = 0
		where id = ?
	`, userId)
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
