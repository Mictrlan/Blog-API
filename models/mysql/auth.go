package mysql

import (
	"database/sql"
)

// AddAuth -
func AddAuth(db *sql.DB, username, password string) (err error) {
	_, err = db.Exec("INSERT INTO auths(username,password)VALUES(?, ?)", username, password)
	return err
}

// ModifyPwd -
func ModifyPwd(db *sql.DB, username, password string) (err error) {
	_, err = db.Exec("UPDATE auths SET password = ? WHERE username = ?", password, username)
	return err
}

// GetAuthID -
func GetAuthID(db *sql.DB, username, password string) (id int, err error) {
	err = db.QueryRow("SELECT id FROM auths WHERE username = ? AND password = ?", username, password).Scan(&id)
	return id, err
}
