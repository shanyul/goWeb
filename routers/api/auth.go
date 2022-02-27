package api

import (
	"designer-api/internal/request"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type UserApi struct {
	userService    service.UserService
	captchaService service.CaptchaService
}

func (api *UserApi) Login(c *gin.Context) {
	var form request.LoginUserForm
	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	checkCaptcha := api.captchaService.Verify(form.CaptchaId, form.Captcha)
	if !checkCaptcha {
		app.Response(c, httpCode, e.ERROR_CHECK_CAPTCHA_FAIL, nil, "")
		return
	}

	loginData := service.User{}
	loginData.Username = form.Username
	loginData.Password = form.Password

	data, code := api.userService.CheckUser(&loginData)

	app.Response(c, http.StatusOK, code, data, "")
}

func (api *UserApi) Register(c *gin.Context) {
	var form request.RegisterUserForm
	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	checkCaptcha := api.captchaService.Verify(form.CaptchaId, form.Captcha)
	if !checkCaptcha {
		app.Response(c, httpCode, e.ERROR_CHECK_CAPTCHA_FAIL, nil, "")
		return
	}

	code := e.INVALID_PARAMS
	if strings.Compare(form.Password, form.ConfirmPassword) != 0 {
		code = e.ERROR_CONFIRM_PASSWORD_NOT_EQ
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}
	if exist, _ := api.userService.ExistNickname(form.Username); exist {
		code = e.ERROR_USER_NICKNAME_EXIST
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.MinCost)
	if err != nil {
		code = e.ERROR_GENERATE_PASSWORD
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}

	authData := service.User{}
	authData.Username = form.Username
	authData.Nickname = form.Nickname
	authData.Password = string(hashPwd)

	if err := api.userService.AddUser(&authData); err != nil {
		code = e.ERROR_REGISTER_FAIL
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

func (api *UserApi) RefreshToken(c *gin.Context) {
	token := c.GetHeader("token")
	code := e.ERROR_TOKEN_NOT_EXIST
	data := make(map[string]interface{})
	if token == "" {
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}

	newToken, err := util.RefreshToken(token)
	if err != nil {
		code = e.ERROR_AUTH
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}
	data["token"] = newToken
	app.Response(c, http.StatusOK, e.SUCCESS, data, "")
}

// 修改用户信息
func (api *UserApi) EditUser(c *gin.Context) {
	var form request.EditUserForm
	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := api.userService.GetUserInfo(id)

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
	userData.Profession = form.Profession
	userData.Charge = form.Charge
	userData.Introduction = form.Introduction

	if userInfo.Username != form.Username {
		if exist, _ := api.userService.ExistNickname(form.Username); exist {
			app.Response(c, http.StatusOK, e.ERROR_USER_NICKNAME_EXIST, nil, "")
			return
		}
	}

	err := api.userService.Edit(userData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

func (api *UserApi) GetUserInfo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	userInfo := api.userService.GetUserInfo(id)
	app.Response(c, http.StatusOK, e.SUCCESS, userInfo, "")
}

func (api *UserApi) ChangePassword(c *gin.Context) {
	var (
		form request.ChangePasswordForm
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	code := e.INVALID_PARAMS
	if strings.Compare(form.Password, form.ConfirmPassword) != 0 {
		code = e.ERROR_CONFIRM_PASSWORD_NOT_EQ
		app.Response(c, http.StatusOK, code, nil, "")
		return
	}
	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	code = api.userService.ChangePassword(id, form.Password, form.OldPassword)
	app.Response(c, http.StatusOK, code, nil, "")
}
