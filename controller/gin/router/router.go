package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

// InitRouter return router
func InitRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	return r
}
