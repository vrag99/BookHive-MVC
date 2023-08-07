package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
	"database/sql"
)

func AddUser(username string, password string, confirmPassword string, regAsAdmin bool, adminPassword string) types.Err {
	db, _ := Connection()

	config, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}
	ADMIN_PASS := config.AdminPassword

	var userId int
	userExists := db.QueryRow("select id from users where username=?", username).Scan(&userId) != sql.ErrNoRows
	defer db.Close()

	var adminApproved bool
	if regAsAdmin && adminPassword == ADMIN_PASS { adminApproved = true } else {adminApproved = false}

	if userExists {
		return types.Err{ErrMsg: "User already exists"}
	} else if password != confirmPassword{
		return types.Err{ErrMsg: "The passwords don't match"}
	} else if regAsAdmin && !adminApproved {
		return types.Err{ErrMsg: "Incorrect admin passcode"}
	} else{
		hashedPassword := utils.HashPassword(password)
		utils.ExecSql(db, "insert into users (username, admin, hash) values(?, ?, ?)", username, regAsAdmin, hashedPassword)
		return types.Err{}
	}
}
