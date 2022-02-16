package api

import (
	"designer-api/internal/request"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type UserApi struct {
	userService service.UserService
}

func (api *UserApi) Login(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.LoginUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	loginData := service.User{}
	loginData.Nickname = form.Nickname
	loginData.Password = form.Password

	data, code := api.userService.CheckUser(&loginData)

	appG.Response(http.StatusOK, code, data)
}

func (api *UserApi) Register(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.RegisterUserForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	code := e.INVALID_PARAMS
	if strings.Compare(form.Password, form.ConfirmPassword) != 0 {
		code = e.ERROR_CONFIRM_PASSWORD_NOT_EQ
		appG.Response(http.StatusOK, code, nil)
		return
	}
	if exist, _ := api.userService.ExistNickname(form.Nickname); exist {
		code = e.ERROR_USER_NICKNAME_EXIST
		appG.Response(http.StatusOK, code, nil)
		return
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.MinCost)
	if err != nil {
		code = e.ERROR_GENERATE_PASSWORD
		appG.Response(http.StatusOK, code, nil)
		return
	}

	authData := service.User{}
	authData.Username = form.Username
	authData.Nickname = form.Nickname
	authData.Password = string(hashPwd)

	if err := api.userService.AddUser(&authData); err != nil {
		code = e.ERROR_REGISTER_FAIL
		appG.Response(http.StatusOK, code, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func (api *UserApi) RefreshToken(c *gin.Context) {
	token := c.GetHeader("token")
	code := e.ERROR_TOKEN_NOT_EXIST
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	if token == "" {
		appG.Response(http.StatusOK, code, nil)
		return
	}

	newToken, err := util.RefreshToken(token)
	if err != nil {
		code = e.ERROR_AUTH
		appG.Response(http.StatusOK, code, nil)
		return
	}
	data["token"] = newToken
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// 修改文章作品
func (api *UserApi) EditUser(c *gin.Context) {
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

	userData := service.User{}
	userData.UserId = id
	userData.Username = form.Username
	userData.Nickname = form.Nickname
	userData.Avatar = form.Avatar
	userData.BgImage = form.BgImage
	userData.Province = form.Province
	userData.City = form.City
	userData.Distinct = form.Distinct
	userData.Address = form.Address
	userData.Remark = form.Remark

	err := api.userService.Edit(userData)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
