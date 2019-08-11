package mysql

import (
	"database/sql"
	"time"
)

// Tag -
type Tag struct {
	ID        int
	Name      string
	CreatedBy string
	CreatedOn int
	UpdatedBy string
	UpdatedOn int
	State     bool
	DeletedOn int
}

// AddTag -
func AddTag(db *sql.DB, name, createdBy string) (err error) {
	_, err = db.Exec("INSERT INTO tags(name,created_by)VALUES(?, ?)", name, createdBy)
	return err
}

// SoftDeleteTag -
func SoftDeleteTag(db *sql.DB, id int) (err error) {
	_, err = db.Exec("UPDATE tags SET state = false, delete_on = ? WHERE id = ?", time.Now().Unix(), id)
	return err
}

// HardDeleteTag -
func HardDeleteTag(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM tags WHERE id = ?", id)
	return err
}

// EditTag -
func EditTag(db *sql.DB, name, updateBy string, id int) (err error) {
	_, err = db.Exec("UPATE tags SET name = ?, updateBy = ? WHERE id = ?", name, updateBy, id)
	return err
}

// GetTagByID -
func GetTagByID(db *sql.DB, id int) (result interface{}, err error) {
	var tags Tag

	if err = db.QueryRow("SELECT FROM tags WHERE id = ?", id).Scan(
		&tags.ID,
		&tags.Name,
		&tags.CreatedBy,
		&tags.CreatedOn,
		&tags.UpdatedBy,
		&tags.UpdatedOn,
		&tags.State,
		&tags.DeletedOn,
	); err != nil {
		return nil, err
	}

	return tags, nil
}

// GetTags -
func GetTags(db *sql.DB) (count int, result []interface{}, err error) {

	var tags Tag
	rows, err := db.Query("SELECT * FROM tags WHERE state = ?", true)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&tags.ID,
			&tags.Name,
			&tags.CreatedBy,
			&tags.CreatedOn,
			&tags.UpdatedBy,
			&tags.UpdatedOn,
			&tags.State,
			&tags.DeletedOn,
		); err != nil {
			return 0, nil, err
		}

		result = append(result, tags)
	}

	return len(result), result, nil
}
