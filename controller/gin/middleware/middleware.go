package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Mictrlan/blog-api/controller/gin/api"
	jwt "github.com/appleboy/gin-jwt"
	gojwt "gopkg.in/dgrijalva/jwt-go.v3"
)

// Auth -
func Auth(authCtl *api.AuthController) *jwt.GinJWTMiddleware {

	return &jwt.GinJWTMiddleware{
		Realm:      "test",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(ctx *gin.Context) (interface{}, error) {
			return authCtl.Login(ctx)
		},

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int); ok {
				if v != 0 {
					return jwt.MapClaims{
						"userID": v,
					}
				}
			}
			return jwt.MapClaims{}
		},

		IdentityHandler: func(mapclaims gojwt.MapClaims) interface{} {
			return mapclaims["userID"]
		},
	}
}
