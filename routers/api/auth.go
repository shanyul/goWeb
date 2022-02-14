package api

import (
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(20)"`
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	appG := app.Gin{C: c}

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.ERROR_LOGIN_PARAMS
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, code, data)
		return
	}

	authService := service.Auth{
		Username: username,
		Password: password,
	}
	data, code = authService.CheckUser()

	appG.Response(http.StatusOK, code, data)
}

func RefreshToken(c *gin.Context) {
	token := c.GetHeader("token")
	code := e.ERROR_TOKEN_NOT_EXIST
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	if token == "" {
		appG.Response(http.StatusOK, code, data)
		return
	}

	newToken, err := util.RefreshToken(token)
	if err != nil {
		code = e.ERROR_AUTH
		appG.Response(http.StatusOK, code, data)
		return
	}
	data["token"] = newToken
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
