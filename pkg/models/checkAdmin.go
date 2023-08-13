package models

func CheckAdmin(userId interface{}) (bool, error) {
	db, _ := Connection()
	defer db.Close()

	var isAdmin bool
	err := db.QueryRow(`select admin from users where id=?`, userId).Scan(&isAdmin)
	if err != nil {
		return false, err
	}

	return isAdmin, nil
}
