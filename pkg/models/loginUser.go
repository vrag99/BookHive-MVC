package models

import (
	"BookHive/pkg/types"
	"BookHive/pkg/utils"
)

func GetJWT(username string, password string) (string, types.Err, bool) {
	// Returns jwt, error, isAdmin
	db, err := Connection()
	if err != nil {
		return "", types.Err{ErrMsg: "Error connecting to db"}, false
	}

	rows, err := utils.ExecSql(db, "select * from users where username=?", username)
	if err != nil {
		return "", types.Err{ErrMsg: "Error fetching users"}, false
	}

	defer rows.Close()
	defer db.Close()

	if !rows.Next() {
		return "", types.Err{ErrMsg: "User doesn't exist"}, false
	} else {
		var userData types.UserData
		err := rows.Scan(&userData.Id, &userData.Username, &userData.Admin, &userData.Hash, &userData.RequestAdmin)
		if err != nil {
			panic(err)
		}

		passMatch := utils.MatchPassword(password, userData.Hash)
		if passMatch {
			token := utils.GenerateJWT(userData)
			isAdmin := userData.Admin != 0
			return token, types.Err{}, isAdmin

		} else {
			return "", types.Err{ErrMsg: "Incorrect password"}, false
		}
	}
}
