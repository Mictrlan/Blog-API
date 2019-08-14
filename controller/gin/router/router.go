package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/Mictrlan/blog-api/controller/gin/api"

	mw "github.com/Mictrlan/blog-api/controller/gin/middleware"
)

// InitRouter return router
func InitRouter(db *sql.DB) *gin.Engine {

	//	f, _ := os.Create("./pkg/log/gin.log")
	//	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()

	authCtl := api.NewAuthCtl(db)
	tagCtl := api.NewTagCtl(db)
	articleCtl := api.NewArticleCtl(db)

	r.POST("api/v1/add/auth", authCtl.AddAuth)

	AuthMiddleware := mw.Auth(authCtl)

	r.POST("api/v1/login", AuthMiddleware.LoginHandler)

	r.Use(func(ctx *gin.Context) {
		AuthMiddleware.MiddlewareFunc()(ctx)
	})

	apiv1 := r.Group("api/v1")
	{
		apiv1.PUT("/modifyPwd", authCtl.ModifyPwd)

		apiv1.POST("/add/tag", tagCtl.AddTag)
		apiv1.DELETE("/delete/tag", tagCtl.DeleteTag)
		apiv1.DELETE("/remove/tag", tagCtl.RemoveTag)
		apiv1.PUT("/update/tag", tagCtl.EditTag)
		apiv1.POST("/get/tag", tagCtl.GetTagByID)
		apiv1.POST("/get/tags", tagCtl.GetTags)

		apiv1.POST("/add/article", articleCtl.AddArticle)
		apiv1.DELETE("/delete/article", articleCtl.DeleteArticle)
		apiv1.DELETE("/remove/article", articleCtl.RemoveArticle)
		apiv1.PUT("/update/article", articleCtl.EditArticle)
		apiv1.POST("/get/article", articleCtl.GetArticleByID)
		apiv1.POST("/get/articles", articleCtl.GetArticleByTagID)
	}

	return r
}
