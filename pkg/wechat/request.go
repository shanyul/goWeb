package wechat

import (
	"designer-api/pkg/gredis"
	"designer-api/pkg/setting"
	"designer-api/pkg/util"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

type Wechat struct{}

const (
	wechatApi               = "https://api.weixin.qq.com"
	wechatOpenApi           = "https://open.weixin.qq.com"
	checkParam              = "appid=%s&secret=%s"
	prefixAccessTokenKey    = "prefix_wechat_access_token"
	prefixWebAccessTokenKey = "prefix_web_access_token"
)

// Code2Session 小程序登录
func (*Wechat) Code2Session(code string) (Code2SessionResponseForm, error) {
	params := fmt.Sprintf(checkParam+"&grant_type=authorization_code&js_code=%s", setting.WechatSetting.AppId, setting.WechatSetting.AppSecret, code)
	requestUrl := wechatApi + "/sns/jscode2session?" + params
	requestData, err := util.Get(requestUrl, 5)
	var response Code2SessionResponseForm
	if err != nil {
		return response, err
	}
	_ = json.Unmarshal([]byte(requestData), &response)

	return response, nil
}

// GetAccessToken 获取小程序 access token
func (*Wechat) GetAccessToken() (AccessTokenResponseForm, error) {
	var response AccessTokenResponseForm
	cacheData, err := gredis.Get(prefixAccessTokenKey)
	if cacheData != nil && err == nil {
		_ = json.Unmarshal(cacheData, &response)
	}

	if response.AccessToken == "" {
		params := fmt.Sprintf(checkParam+"&grant_type=client_credential", setting.WechatSetting.AppId, setting.WechatSetting.AppSecret)
		requestUrl := wechatApi + "/cgi-bin/token?" + params

		requestData, err := util.Get(requestUrl, 5)
		if err != nil {
			return response, err
		}
		_ = json.Unmarshal([]byte(requestData), &response)

		if response.AccessToken != "" {
			_ = gredis.Set(prefixAccessTokenKey, response, response.ExpiresIn)
		}
	}

	return response, nil
}

// GetWebCode 微信扫码登录地址
func (*Wechat) GetWebCode() string {
	redirectUrl := setting.AppSetting.AppHost + "/wechat/web-callback?timestamp=" + strconv.FormatInt(time.Now().Unix(), 10)
	sign := util.StringToMd5(redirectUrl)
	urlEncode := url.QueryEscape(redirectUrl)
	requestCodeUrl := fmt.Sprintf("%s/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_login&state=%s#wechat_redirect",
		wechatOpenApi, setting.WechatSetting.WebAppId, urlEncode, sign)

	return requestCodeUrl
}

// GetWebAccessToken 获取网页端微信 access token
func (*Wechat) GetWebAccessToken(code string) (WebAccessTokenResponseForm, error) {
	var response WebAccessTokenResponseForm
	cacheData, err := gredis.Get(prefixWebAccessTokenKey)
	if cacheData != nil && err == nil {
		_ = json.Unmarshal(cacheData, &response)
	}

	if response.AccessToken == "" {
		params := fmt.Sprintf(checkParam+"&code=%s&grant_type=authorization_code", setting.WechatSetting.WebAppId, setting.WechatSetting.WebAppSecret, code)
		requestUrl := wechatApi + "/sns/oauth2/access_token?" + params

		requestData, err := util.Get(requestUrl, 5)
		if err != nil {
			return response, err
		}
		_ = json.Unmarshal([]byte(requestData), &response)

		if response.AccessToken != "" {
			_ = gredis.Set(prefixAccessTokenKey, response, response.ExpiresIn)
		}
	}

	return response, nil
}

func (*Wechat) GetUserInfo(token string, openId string) (UserInfoResponseForm, error) {
	var response UserInfoResponseForm

	params := fmt.Sprintf("access_token=%s&openid=%s", token, openId)
	requestUrl := wechatApi + "/sns/userinfo?" + params

	requestData, err := util.Get(requestUrl, 5)
	if err != nil {
		return response, err
	}
	_ = json.Unmarshal([]byte(requestData), &response)

	return response, nil
}
