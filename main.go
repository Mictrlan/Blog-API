package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mictrlan/blog-api/controller/gin/router"
	"github.com/Mictrlan/blog-api/models/mysql"
	"github.com/facebookgo/grace/gracehttp"
)

func main() {
	fmt.Println("Hollow World!")

	log.SetFlags(log.Ldate | log.Lshortfile)

	db := mysql.InitModel()

	r := router.InitRouter(db)

	fmt.Println("pid is: ", os.Getpid())

	gracehttp.Serve(&http.Server{
		Addr:    ":8080",
		Handler: r,
	})
}
