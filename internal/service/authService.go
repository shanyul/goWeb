package service

import (
	"designer-api/internal/models"
	"designer-api/pkg/e"
	"designer-api/pkg/gredis"
	"designer-api/pkg/util"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
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

func (service *UserService) GetUserInfo(id int) User {
	key := fmt.Sprintf("%s:%d", prefixLoginKey, id)
	cacheData, _ := gredis.Get(key)
	var userInfo User
	_ = json.Unmarshal(cacheData, &userInfo)

	return userInfo
}

func (service *UserService) ExistNickname(nickname string) (bool, error) {
	return service.UserModel.ExistNickname(nickname)
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
	authInfo, err := service.UserModel.GetByNickname(a.Nickname)
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

	token, err := util.GenerateToken(authInfo.UserId)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
	} else {
		info = make(map[string]interface{})
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
		info["token"] = token

		// 保存用户信息
		key := fmt.Sprintf("%s:%d", prefixLoginKey, authInfo.UserId)
		ttl := 60 * 60 * 24
		_ = gredis.Set(key, info, ttl)

		code = e.SUCCESS
	}

	return
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

	return nil
}
