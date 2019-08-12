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
	CreatedOn string
	UpdatedBy string
	UpdatedOn string
	State     bool
	DeletedOn int
}

// AddArticle -
func AddArticle(db *sql.DB, tagID int, title, description, content, createdBy string) (err error) {
	_, err = db.Exec("INSERT INTO articles(tag_id,title,description,content,created_by)VALUES(?,?,?,?,?)",
		tagID, title, description, content, createdBy)
	return err
}

// SoftDeleteArticle -
func SoftDeleteArticle(db *sql.DB, id int) (err error) {
	_, err = db.Exec("UPDATE articles SET state = false, deleted_on = ? WHERE id = ?", time.Now().Unix(), id)
	return err
}

// HardDeleteArticle -
func HardDeleteArticle(db *sql.DB, id int) (err error) {
	_, err = db.Exec("DELETE FROM articles WHERE id = ?", id)
	return err
}

// EditArticle -
func EditArticle(db *sql.DB, id, tagID int, title, description, content, updatedBy string) (err error) {
	_, err = db.Exec("UPDATE articles SET tag_id = ?,title = ?, description = ?, content = ?, updated_by = ? WHERE id = ?",
		tagID, title, description, content, updatedBy, id)
	return err
}

// GetArticleByID -
func GetArticleByID(db *sql.DB, id int) (*Article, *Tag, error) {
	var (
		articles Article
		tags     Tag
	)
	if err := db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(
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
		return nil, nil, err
	}

	if articles.TagID != 0 {
		if err := db.QueryRow("SELECT * FROM tags WHERE id = ?", articles.TagID).Scan(
			&tags.ID,
			&tags.Name,
			&tags.CreatedBy,
			&tags.CreatedOn,
			&tags.UpdatedBy,
			&tags.UpdatedOn,
			&tags.State,
			&tags.DeletedOn,
		); err != nil {
			return nil, nil, err
		}
	}

	return &articles, &tags, nil
}

// GetArticlesByTag -
func GetArticlesByTag(db *sql.DB, tagID int) (count int, articles []*Article, err error) {
	var article Article

	rows, err := db.Query("SELECT * FROM articles WHERE tag_id = ?", tagID)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	for rows.Next() {

		if err = rows.Scan(
			&article.ID,
			&article.TagID,
			&article.Title,
			&article.Desc,
			&article.Content,
			&article.CreatedBy,
			&article.CreatedOn,
			&article.UpdatedBy,
			&article.UpdatedOn,
			&article.State,
			&article.DeletedOn,
		); err != nil {
			return 0, nil, err
		}

		articles = append(articles, &article)
	}

	return len(articles), articles, nil
}
