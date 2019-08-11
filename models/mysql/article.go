package mysql

import (
	"database/sql"
	"time"
)

// Article -
type Article struct {
	ID        int
	TagID     int
	Title     string
	Desc      string
	Content   string
	CreatedBy string
	CreatedOn int
	UpdatedBy string
	UpdatedOn int
	State     bool
	DeletedOn int
}

// AddArticle -
func AddArticle(db *sql.DB, tagID int, title, description, content, createdBy string) (err error) {
	_, err = db.Exec("INSERT INTO articles(tagID,title,description,content,created_by)VALUES(?,?,?,?,?)",
		tagID, title, description, content, createdBy)
	return err
}

// SoftDeleteArticle -
func SoftDeleteArticle(db *sql.DB, id int) (err error) {
	_, err = db.Exec("UPDATE articles SET state = false, delete_on = ? WHERE id = ?", time.Now().Unix(), id)
	return err
}

// HardDeleteArticle -
func HardDeleteArticle(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

// EditArticle -
func EditArticle(db *sql.DB, id, tagID int, title, description, content, updatedBy string) (err error) {
	_, err = db.Exec("UPDATE articles SET tag_id = ?,title = ?, descroption = ?, content = ?, updated_by = ?",
		id, tagID, title, description, content, updatedBy)
	return err
}

// GetArticleByID -
func GetArticleByID(db *sql.DB, id int) (result []interface{}, err error) {
	var articles Article
	var tags Tag

	if err = db.QueryRow("SELECT * FROM articles,tags WHERE id = ?", id).Scan(
		&articles.ID,
		&articles.TagID,
		&articles.Title,
		&articles.Desc,
		&articles.Content,
		&articles.CreatedBy,
		&articles.CreatedOn,
		&articles.UpdatedBy,
		&articles.UpdatedOn,
		&articles.State,
		&articles.DeletedOn,
	); err != nil {
		return nil, err
	}

	if articles.TagID != 0 {
		if err = db.QueryRow("SELECT FROM tags WHERE id = ?", articles.TagID).Scan(
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
	}

	result = append(result, articles, tags)

	return result, nil
}

// GetArticlesByTag -
func GetArticlesByTag(db *sql.DB, tagID int) (count int, result []interface{}, err error) {
	var articles Article

	rows, err := db.Query("SELECT * FROM articles WHERE tag_id = ?", tagID)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {

		if err = rows.Scan(
			&articles.ID,
			&articles.TagID,
			&articles.Title,
			&articles.Desc,
			&articles.Content,
			&articles.CreatedBy,
			&articles.CreatedOn,
			&articles.UpdatedBy,
			&articles.UpdatedOn,
			&articles.State,
			&articles.DeletedOn,
		); err != nil {
			return 0, nil, err
		}

		result = append(result, articles)
	}

	return len(result), result, nil
}
