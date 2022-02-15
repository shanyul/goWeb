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

type Auth struct {
	Username        string
	Nickname        string
	Password        string
	ConfirmPassword string
}

type UserInfo struct {
	UserId     int
	Username   string
	Nickname   string
	Avatar     string
	BgImage    string
	Phone      string
	Email      string
	State      int
	Province   string
	City       string
	Distinct   string
	Address    string
	CreateTime string
	Token      string
}

const prefixLoginKey = "key_user_login"

func GetUserInfo(id int) UserInfo {
	key := fmt.Sprintf("%s:%d", prefixLoginKey, id)
	cacheData, _ := gredis.Get(key)
	var userInfo UserInfo
	_ = json.Unmarshal(cacheData, &userInfo)

	return userInfo
}

func (a *Auth) ExistByName() (bool, error) {
	return models.ExistNickname(a.Nickname)
}

func (a *Auth) AddUser() error {
	category := map[string]interface{}{
		"username": a.Username,
		"nickname": a.Nickname,
		"password": a.Password,
	}

	if err := models.AddUser(category); err != nil {
		return err
	}

	return nil
}

func (a *Auth) CheckUser() (info map[string]interface{}, code int) {
	authInfo, err := models.GetByNickname(a.Nickname)
	code = e.SUCCESS
	if err != nil {
		code = e.ERROR_LOGIN_PARAMS
		return
	}
	// 验证密码是否正确
	fmt.Println(authInfo.Password, a.Password)
	err = bcrypt.CompareHashAndPassword([]byte(authInfo.Password), []byte(a.Password))
	fmt.Println("errrr", err)
	if err != nil {
		code = e.ERROR_LOGIN_PARAMS
		return
	}

	token, err := util.GenerateToken(authInfo.UserID)
	if err != nil {
		code = e.ERROR_AUTH_TOKEN
	} else {
		info = make(map[string]interface{})
		info["userId"] = authInfo.UserID
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
		key := fmt.Sprintf("%s:%d", prefixLoginKey, authInfo.UserID)
		ttl := 60 * 60 * 24
		_ = gredis.Set(key, info, ttl)

		code = e.SUCCESS
	}

	return
}
