package service

import (
	"designer-api/internal/models"
	"designer-api/pkg/e"
	"designer-api/pkg/gredis"
	"designer-api/pkg/util"
	"encoding/json"
	"fmt"
)

type Auth struct {
	Username string
	Password string
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

func (a *Auth) CheckUser() (info map[string]interface{}, code int) {
	authInfo, result := models.CheckAuth(a.Username, a.Password)
	code = e.SUCCESS
	if !result {
		code = e.ERROR_LOGIN_PARAMS
		return info, code
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
		gredis.Set(key, info, ttl)

		code = e.SUCCESS
	}

	return info, code
}
