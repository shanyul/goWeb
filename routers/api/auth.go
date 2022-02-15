package api

import (
	"designer-api/internal/models"
	"designer-api/internal/request"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type auth struct {
	Username        string `valid:"Required; MaxSize(20)"`
	Nickname        string `valid:"Required; AlphaNumeric; MaxSize(20)"`
	Password        string `valid:"Required; MaxSize(20)"`
	ConfirmPassword string `valid:"Required; MaxSize(20)"`
}

type LoginAuth struct {
	Nickname string `valid:"Required; AlphaNumeric; MaxSize(20)"`
	Password string `valid:"Required; MaxSize(20)"`
}

func Login(c *gin.Context) {
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	appG := app.Gin{C: c}

	valid := validation.Validation{}
	a := LoginAuth{Nickname: nickname, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.ERROR_LOGIN_PARAMS
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, code, data)
		return
	}

	authService := service.Auth{
		Nickname: nickname,
		Password: password,
	}
	data, code = authService.CheckUser()

	appG.Response(http.StatusOK, code, data)
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	nickname := c.PostForm("nickname")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirmPassword")
	appG := app.Gin{C: c}

	valid := validation.Validation{}
	a := auth{Username: username, Nickname: nickname, Password: password, ConfirmPassword: confirmPassword}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, code, data)
		return
	}
	if strings.Compare(a.Password, a.ConfirmPassword) != 0 {
		code = e.ERROR_CONFIRM_PASSWORD_NOT_EQ
		appG.Response(http.StatusOK, code, data)
		return
	}
	if exist, _ := models.ExistNickname(a.Nickname); exist {
		code = e.ERROR_USER_NICKNAME_EXIST
		appG.Response(http.StatusOK, code, data)
		return
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		code = e.ERROR_GENERATE_PASSWORD
		appG.Response(http.StatusOK, code, data)
		return
	}

	authService := service.Auth{
		Username: username,
		Nickname: nickname,
		Password: string(hashPwd),
	}
	if err := authService.AddUser(); err != nil {
		code = e.ERROR_REGISTER_FAIL
		appG.Response(http.StatusOK, code, data)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)
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

// 修改文章作品
func EditUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.EditUserForm{}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userData := service.GetUserInfo(id)

	userService := service.UserInfo{
		UserId:   userData.UserId,
		Username: form.Username,
		Nickname: form.Nickname,
		Avatar:   form.Avatar,
		BgImage:  form.BgImage,
		Province: form.Province,
		City:     form.City,
		Distinct: form.Distinct,
		Remark:   form.Remark,
	}

	err := userService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
