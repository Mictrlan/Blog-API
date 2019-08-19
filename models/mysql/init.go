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

	// If you want to deploy your service to dokcer, you shoueld replace  127.0.0.1 with docker.for.mac.localhost
	// CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o  executable .  // cmd["executable"]
	// docker build -t servser_name
	// docker run -p 8083:8083 servser_name
	db, err := sql.Open("mysql", "root:Miufighting.@tcp(127.0.0.1)/blog")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range constvar.TablesSQLString {
		db.Exec(v)
	}

	return db
}
