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

// GetAuthIDAndPwd -
func GetAuthIDAndPwd(db *sql.DB, username string) (id int, password string, err error) {
	err = db.QueryRow("SELECT id, password FROM auths WHERE username = ? ", username).Scan(&id, &password)
	return id, password, err
}
