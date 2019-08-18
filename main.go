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

// @title blog-api
// @version 1.0
// @description This is a simple trying RESTful-API
// @termsOfService https://github.com/Mictrlan/blog-api
// @license.name MIT
// @license.url https://github.com/Mictrlan/blog-api/blob/master/LICENSE
func main() {
	fmt.Println("Hollow World!")

	log.SetFlags(log.Ldate | log.Lshortfile)

	db := mysql.InitModel()
	r := router.InitRouter(db)

	// https
	go func() {
		if err := http.ListenAndServeTLS(":8081", "./conf/server.crt", "./conf/server.key", r); err != nil {
			log.Fatal(err)
		}
	}()

	//Graceful Restart
	fmt.Println("pid is: ", os.Getpid())

	gracehttp.Serve(&http.Server{
		Addr:    ":8080",
		Handler: r,
	})

}
