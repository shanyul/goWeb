package middleware

import (
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/logging"
	"designer-api/pkg/setting"
	"designer-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"sort"
	"strings"
	"time"
)

func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 本地环境不需要验签
		if setting.AppSetting.Environment != "local" {
			// 获取参数
			query := c.Request.URL.Query()
			var keys []string
			reqSign := ""
			params := make(map[string]string)
			for key, item := range query {
				if key == "sign" {
					reqSign = item[0]
				} else {
					keys = append(keys, key)
					params[key] = item[0]
				}
			}

			if reqSign == "" {
				app.Response(c, http.StatusUnauthorized, e.INVALID_PARAMS, nil, "")
				c.Abort()
				return
			}

			sort.Strings(keys)
			signString := ""
			for _, key := range keys {
				signString += key + ":" + params[key] + "&"
			}
			signString += setting.AppSetting.SignKey
			checkSign := util.EncodeMD5(signString)
			if strings.Compare(checkSign, reqSign) != 0 {
				logging.Error("checkSignError:", checkSign)
				app.Response(c, http.StatusBadRequest, e.ERROR_AUTH_CHECK_SIGN_FAIL, nil, "")
				c.Abort()
				return
			}

			nowTime := time.Now().Unix()
			expireTime := com.StrTo(params["timestamp"]).MustInt64() + (60 * 10)
			if expireTime < nowTime {
				app.Response(c, http.StatusBadRequest, e.ERROR_AUTH_CHECK_SIGN_EXPIRE, nil, "")
				c.Abort()
				return
			}
		}

		userId := 0
		// 判断是否登录
		token := c.GetHeader("token")
		if token != "" {
			claims, err := util.ParseToken(token)
			if err == nil {
				userId = claims.UsesId
			}
		}

		// 设置默认值
		c.Set("userId", userId)

		c.Next()
	}
}
