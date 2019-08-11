package main

import (
	"fmt"
	"log"

	"github.com/Mictrlan/blog-api/controller/gin/router"
	"github.com/Mictrlan/blog-api/models/mysql"
)

func main() {
	fmt.Println("Hollow World!")

	log.SetFlags(log.Ldate | log.Lshortfile)

	db := mysql.InitModel()

	r := router.InitRouter(db)

	r.Run()
}
