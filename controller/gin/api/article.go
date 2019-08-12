package api

import (
	"database/sql"
	"net/http"

	"github.com/Mictrlan/blog-api/models/mysql"
	"github.com/gin-gonic/gin"
)

// ArticleController -
type ArticleController struct {
	db *sql.DB
}

// NewArticleCtl -
func NewArticleCtl(db *sql.DB) *ArticleController {
	return &ArticleController{
		db: db,
	}
}

// AddArticle -
func (articleCtl *ArticleController) AddArticle(ctx *gin.Context) {
	var (
		req struct {
			TagID     int    `json:"tag_id" binding:"gte=0"`
			Title     string `json:"title" binding:"required"`
			Desc      string `json:"description"`
			Content   string `json:"content"`
			CreatedBy string `json:"created_by" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.AddArticle(articleCtl.db, req.TagID, req.Title, req.Desc, req.Content, req.CreatedBy); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// DeleteArticle -
func (articleCtl *ArticleController) DeleteArticle(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.SoftDeleteArticle(articleCtl.db, req.ID); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// RemoveArticle -
func (articleCtl *ArticleController) RemoveArticle(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" binding:"gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.HardDeleteArticle(articleCtl.db, req.ID); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// EditArticle -
func (articleCtl *ArticleController) EditArticle(ctx *gin.Context) {
	var (
		req struct {
			ID        int    `json:"id" binding:"required,gt=0"`
			TagID     int    `json:"tag_id" binding:"gte=0"`
			Title     string `json:"title"`
			Desc      string `json:"description"`
			Content   string `json:"content"`
			UpdatedBy string `json:"updated_by" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	original, _, err := mysql.GetArticleByID(articleCtl.db, req.ID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	if req.TagID == 0 {
		req.TagID = original.TagID
	}

	if req.Title == "" {
		req.Title = original.Title
	}

	if req.Desc == "" {
		req.Desc = original.Desc
	}

	if req.Content == "" {
		req.Content = original.Content
	}

	if err = mysql.EditArticle(articleCtl.db, req.ID, req.TagID, req.Title, req.Desc, req.Content, req.UpdatedBy); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// GetArticleByID -
func (articleCtl *ArticleController) GetArticleByID(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" bindig:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	article, tag, err := mysql.GetArticleByID(articleCtl.db, req.ID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"article": article,
		"tag":     tag,
	})
}

// GetArticleByTagID -
func (articleCtl *ArticleController) GetArticleByTagID(ctx *gin.Context) {
	var (
		req struct {
			TagID int `json:"tag_id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, articles, err := mysql.GetArticlesByTag(articleCtl.db, req.TagID)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"count":    count,
		"articles": articles,
	})
}
