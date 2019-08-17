package mysql

import (
	"database/sql"
	"time"
)

// Insert  add file info to table
func Insert(db *sql.DB, articleID int, path, md5 string) error {
	result, err := db.Exec("INSERT INTO files(article_id,md5,path,created_at) VALUES (?,?,?,?)", articleID, path, md5, time.Now())
	if err != nil {
		return err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return err
	}

	return nil
}

// QueryPathByMD5 query path by md5
func QueryPathByMD5(db *sql.DB, md5 string) (string, error) {
	var path string

	err := db.QueryRow("SELECT path FROM files WHERE md5 = ? LOCK IN SHARE MODE", md5).Scan(&path)

	return path, err
}
