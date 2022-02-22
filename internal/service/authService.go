package service

import (
	"designer-api/internal/models"
	"designer-api/pkg/e"
	"designer-api/pkg/gredis"
	"designer-api/pkg/util"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	UserModel models.UserModel
}

type User struct {
	models.User
	Token           string
	ConfirmPassword string
}

const prefixLoginKey = "key_user_login"

func (service *UserService) GetUserInfo(id int) (userInfo models.User) {
	key := fmt.Sprintf("%s:%d", prefixLoginKey, id)
	cacheData, err := gredis.Get(key)
	if cacheData != nil && err == nil {
		_ = json.Unmarshal(cacheData, &userInfo)
	}

	// 缓存不存在取数据库
	if userInfo.UserId == 0 {
		userInfo, _ = service.UserModel.GetByUserId(id)
		_, _ = service.saveUser(userInfo)
	}
	userInfo.Password = "*"

	return userInfo
}

func (service *UserService) ExistNickname(username string) (bool, error) {
	return service.UserModel.ExistNickname(username)
}

func (service *UserService) GetUserById(id int) models.User {
	userInfo, _ := service.UserModel.GetByUserId(id)
	return userInfo
}

func (service *UserService) AddUser(a *User) error {
	user := models.User{}
	user.Username = a.Username
	user.Nickname = a.Nickname
	user.Password = a.Password

	if err := service.UserModel.AddUser(&user); err != nil {
		return err
	}

	return nil
}

func (service *UserService) CheckUser(a *User) (info map[string]interface{}, code int) {
	authInfo, err := service.UserModel.GetByNickname(a.Username)
	code = e.SUCCESS
	if err != nil {
		code = e.ERROR_LOGIN_PARAMS
		return
	}
	// 验证密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(a.Password))
	if err != nil {
		code = e.ERROR_LOGIN_PARAMS
		return
	}

	token, err := util.GenerateToken(authInfo.UserId, 0)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
		return
	}

	info, err = service.saveUser(authInfo)
	if err != nil {
		code = e.ERROR_LOGIN_FAIL
		return
	}
	info["token"] = token

	return
}

func (service *UserService) saveUser(authInfo models.User) (map[string]interface{}, error) {
	info := make(map[string]interface{})

	info["userId"] = authInfo.UserId
	info["username"] = authInfo.Username
	info["nickname"] = authInfo.Nickname
	info["avatar"] = authInfo.Avatar
	info["bgImage"] = authInfo.BgImage
	info["phone"] = authInfo.Phone
	info["email"] = authInfo.Email
	info["state"] = authInfo.State
	info["province"] = authInfo.Province
	info["city"] = authInfo.City
	info["distinct"] = authInfo.Distinct
	info["address"] = authInfo.Address
	info["createTime"] = authInfo.CreateTime

	// 保存用户信息
	key := fmt.Sprintf("%s:%d", prefixLoginKey, authInfo.UserId)
	ttl := 60 * 60 * 6
	err := gredis.Set(key, info, ttl)

	if err != nil {
		return nil, err
	}
	return info, nil
}

func (service *UserService) Edit(u User) error {
	user := models.User{}
	user.Username = u.Username
	user.Nickname = u.Nickname
	user.Avatar = u.Avatar
	user.BgImage = u.BgImage
	user.Province = u.Province
	user.City = u.City
	user.Distinct = u.Distinct
	user.Address = u.Address
	user.Remark = u.Remark

	if err := service.UserModel.EditUser(u.UserId, user); err != nil {
		return err
	}

	// 更新缓存信息
	userInfo, _ := service.UserModel.GetByUserId(u.UserId)
	_, _ = service.saveUser(userInfo)

	return nil
}

// CheckByCode 检查用户通过 openid
func (service *UserService) CheckByCode(code string, sessionKey string, unionId string) (info map[string]interface{}, responseCode int) {
	authInfo, err := service.UserModel.GetByCode(code)
	responseCode = e.SUCCESS
	if err != nil {
		responseCode = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		return
	}
	user := models.User{}
	if authInfo.UserId == 0 {
		user.WechatOpenid = code
		user.SessionKey = sessionKey
		user.UnionId = unionId
		authInfo.UserId = service.UserModel.AddWechatUser(&user)
	} else if sessionKey != "" {
		user.SessionKey = sessionKey
		_ = service.UserModel.EditUser(user.UserId, user)
	}

	// 小程序登录报错一个月过期时间
	ttl := 7 * 24 * time.Hour
	token, err := util.GenerateToken(authInfo.UserId, ttl)
	if err != nil {
		responseCode = e.ERROR_AUTH_TOKEN
		return
	}

	info, err = service.saveUser(authInfo)
	if err != nil {
		responseCode = e.ERROR_LOGIN_FAIL
		return
	}
	info["token"] = token

	return
}

func (service *UserService) ChangePassword(userId int, password string, oldPassword string) (code int) {
	code = e.SUCCESS
	userInfo := service.GetUserById(userId)
	if userInfo.UserId == 0 {
		code = e.ERROR_LOGIN_FAIL
		return
	}

	if userInfo.Password != "" {
		// 验证密码是否正确
		err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(oldPassword))
		if err != nil {
			code = e.ERROR_OLD_PASSWORD_FAIL
			return
		}
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		code = e.ERROR_GENERATE_PASSWORD
		return
	}

	authData := models.User{}
	authData.Password = string(hashPwd)

	if err := service.UserModel.EditUser(userId, authData); err != nil {
		code = e.ERROR_PASSWORD_CHANGE_FAIL
		return
	}

	return
}
