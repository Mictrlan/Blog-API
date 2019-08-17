package api

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/Mictrlan/blog-api/models/mysql"
	"github.com/Mictrlan/blog-api/pkg/upload"
	"github.com/gin-gonic/gin"
)

// UploadController -
type UploadController struct {
	db  *sql.DB
	URL string
}

// NewUploandCtl -
func NewUploandCtl(db *sql.DB, URL string) *UploadController {
	return &UploadController{
		db:  db,
		URL: URL,
	}
}

// Upload -
func (uploadCtl *UploadController) Upload(ctx *gin.Context) {
	if ctx.Request.Method != "POST" {
		ctx.JSON(http.StatusMethodNotAllowed, gin.H{"status": http.StatusMethodNotAllowed})
		return
	}

	var (
		req struct {
			ArticleID int `json:"article_id" binding:"required,gt=0"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	defer func() {
		file.Close()
		ctx.Request.MultipartForm.RemoveAll()
	}()
	if err != nil {
		ctx.AbortWithError(http.StatusForbidden, err)
		return
	}

	fileNew, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.AbortWithError(http.StatusNoContent, err)
		return
	}

	md5, err := upload.MD5(fileNew)
	if err != nil {
		ctx.AbortWithError(http.StatusNotImplemented, err)
	}

	filePath, err := mysql.QueryPathByMD5(uploadCtl.db, md5)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":   http.StatusOK,
			"filePath": uploadCtl.URL + filePath,
		})

		return
	}

	fileSuffix := path.Ext(header.Filename)
	filePath = upload.FileUploadDir + "/" + upload.ClassifyBySuffix(fileSuffix) + "/" + md5 + fileSuffix

	if err = upload.CopyFile(filePath, fileNew); err != nil {
		ctx.AbortWithError(http.StatusPreconditionRequired, err)
		return
	}

	if err = mysql.Insert(uploadCtl.db, req.ArticleID, filePath, md5); err != nil {
		ctx.AbortWithError(http.StatusNotModified, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"filePath": uploadCtl.URL + filePath,
	})
}
