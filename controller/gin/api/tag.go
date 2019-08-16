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
// @Summary Add a new tag to the store
// @Description Add a new tag
// @Produce json
// @Param name query string true "Name"
// @Param created_by query string true "CreatedBy"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/add/tag [post]
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
// @Summary softdelete tag by id
// @Description delete tag by id
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/delete/tag [delete]
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
// @Summary harddelete tag by id
// @Description remove tag by id
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/remove/tag [delete]
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
// @Summary update tag by id
// @Description update tag by id
// @Produce json
// @Param id path int true "ID"
// @Param name query string true "Name"
// @Param updated_by query string true "UpdateBy"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/update/tag [put]
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
// @Summary query an tag information by id
// @Description tag tag by id
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router  /api/v1/get/tag [post]
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
// @Summary query tags information
// @Description get tags information
// @Produce json
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router  /api/v1/get/tags [get]
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
