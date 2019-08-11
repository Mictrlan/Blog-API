package mysql

import (
	"database/sql"
	"log"

	"github.com/Mictrlan/blog-api/pkg/constvar"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// InitModel return db
func InitModel() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1:3306)/blog")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range constvar.TablesSQLString {
		db.Exec(v)
	}

	return db
}
