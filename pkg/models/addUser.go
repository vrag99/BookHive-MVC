package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func AddUser(db *sql.DB, username string, password string, confirmPassword string, requestAdmin bool) types.Err {
	var userId int
	userExists := db.QueryRow("select id from users where username=?", username).Scan(&userId) != sql.ErrNoRows

	if userExists {
		return types.Err{ErrMsg: "User already exists"}

	} else if password != confirmPassword{
		return types.Err{ErrMsg: "The passwords don't match"}
		
	} else{
		hashedPassword := utils.HashPassword(password)
		utils.ExecSql(db, "insert into users (username, admin, hash, requestAdmin) values(?, ?, ?, ?)", username, 0, hashedPassword, requestAdmin)
		return types.Err{}
	}
}
