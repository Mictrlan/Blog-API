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
// @Summary Add a new user
// @Description Add a new user for verification
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/add/auth [post]
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
// @Summary Modify user password
// @Description Modify user password
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "PwdNew"
// @Param confirm query string true "Confirm"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/modifyPwd [put]
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

// Login -
// @Summary User login
// @Description User login
// @Produce json
// @Param username query string true "username"
// @Param password query string true "password"
// @Success 200 {string} json "{"stauts":200,"message":"OK"}"
// @Router /api/v1/login [post]
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
