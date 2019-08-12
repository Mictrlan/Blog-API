package api

import (
	"database/sql"
	"net/http"

	"github.com/Mictrlan/blog-api/models/mysql"
	"github.com/Mictrlan/blog-api/pkg/errno"
	"github.com/Mictrlan/blog-api/pkg/util"
	"github.com/gin-gonic/gin"
)

// AuthController -
type AuthController struct {
	db *sql.DB
}

// NewAuthCtl -
func NewAuthCtl(db *sql.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

// AddAuth -
func (authCtl *AuthController) AddAuth(ctx *gin.Context) {
	var (
		req struct {
			Username string `json:"username" binding:"required"`
			Passwrod string `json:"password" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	pwd, err := util.SaltHashGenerate(req.Passwrod)
	if err != nil {
		ctx.AbortWithError(http.StatusTooEarly, err)
		return
	}

	if err = mysql.AddAuth(authCtl.db, req.Username, pwd); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// ModifyPwd -
func (authCtl *AuthController) ModifyPwd(ctx *gin.Context) {
	var (
		req struct {
			Username string `json:"username" binding:"required"`
			PwdNew   string `json:"password" binding:"required"`
			Confirm  string `json:"confirm" binding:"required"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if req.PwdNew != req.Confirm {
		ctx.AbortWithError(http.StatusForbidden, errno.ErrInconsistentPwd)
		return
	}

	newPwd, err := util.SaltHashGenerate(req.PwdNew)
	if err != nil {
		ctx.AbortWithError(http.StatusTooEarly, err)
		return
	}

	if err = mysql.ModifyPwd(authCtl.db, req.Username, newPwd); err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

// Login - id != zero
func (authCtl *AuthController) Login(ctx *gin.Context) (id int, err error) {
	var (
		req struct {
			Username string `json:"username"`
			Passwrod string `json:"password"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		return 0, err
	}

	id, pwd, err := mysql.GetAuthIDAndPwd(authCtl.db, req.Username)
	if err != nil {
		return 0, err
	}

	if !util.SaltHashCompare([]byte(pwd), req.Passwrod) {
		return 0, err
	}

	return id, nil
}
