package api

import (
	"database/sql"
	"net/http"

	"github.com/Mictrlan/blog-api/models/mysql"
	"github.com/gin-gonic/gin"
)

// TagController -
type TagController struct {
	db *sql.DB
}

// NewTagCtl -
func NewTagCtl(db *sql.DB) *TagController {
	return &TagController{
		db: db,
	}
}

// AddTag -
func (tagCtl *TagController) AddTag(ctx *gin.Context) {
	var (
		req struct {
			Name      string `json:"name" binding:"required"`
			CreatedBy string `json:"created_by" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.AddTag(tagCtl.db, req.Name, req.CreatedBy); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// DeleteTag -
func (tagCtl *TagController) DeleteTag(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.SoftDeleteTag(tagCtl.db, req.ID); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// RemoveTag -
func (tagCtl *TagController) RemoveTag(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.HardDeleteTag(tagCtl.db, req.ID); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// EditTag -
func (tagCtl *TagController) EditTag(ctx *gin.Context) {
	var (
		req struct {
			ID        int    `json:"id" binding:"required,gt=0"`
			Name      string `json:"name" binding:"required"`
			UpdatedBy string `json:"updated_by" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := mysql.EditTag(tagCtl.db, req.Name, req.UpdatedBy, req.ID); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// GetTagByID -
func (tagCtl *TagController) GetTagByID(ctx *gin.Context) {
	var (
		req struct {
			ID int `json:"id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	result, err := mysql.GetTagByID(tagCtl.db, req.ID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"result": result,
	})
}

// GetTags -
func (tagCtl *TagController) GetTags(ctx *gin.Context) {
	count, result, err := mysql.GetTags(tagCtl.db)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"count":  count,
		"result": result,
	})
}
