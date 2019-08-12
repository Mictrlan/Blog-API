package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mictrlan/blog-api/controller/gin/router"
	"github.com/Mictrlan/blog-api/models/mysql"
)

func main() {
	fmt.Println("Hollow World!")

	log.SetFlags(log.Ldate | log.Lshortfile)

	db := mysql.InitModel()

	r := router.InitRouter(db)

	fmt.Println("pid is: ", os.Getpid())

	r.Run()
}
