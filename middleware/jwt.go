package middleware

import (
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			data   interface{}
			code   = e.SUCCESS
			userId = 0
			token  = c.GetHeader("token")
		)

		if token == "" {
			code = e.ERROR_TOKEN_NOT_EXIST
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}

			if code == e.SUCCESS {
				userId = claims.UsesId
			}
		}

		if code != e.SUCCESS || userId == 0 {
			app.Response(c, http.StatusUnauthorized, code, data, "")
			c.Abort()
			return
		}
		// 设置登录用户Id
		c.Set("userId", userId)

		c.Next()
	}
}

func Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.GetHeader("token")
		if token == "" {
			code = e.ERROR_TOKEN_NOT_EXIST
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			app.Response(c, http.StatusUnauthorized, code, data, "")
			c.Abort()
			return
		}

		c.Next()
	}
}
