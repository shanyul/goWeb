package api

import (
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/setting"
	"designer-api/pkg/util"
	"designer-api/pkg/wechat"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type WechatApi struct {
	userService service.UserService
	wechat      wechat.Wechat
}

func (api *WechatApi) Login(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil, "")
		return
	}
	response, err := api.wechat.Code2Session(code)
	if err != nil {
		app.Response(c, http.StatusBadRequest, e.ERROR_WECHAT_REQUEST_FAIL, nil, "")
		return
	}

	if response.ErrCode == 40029 {
		app.Response(c, http.StatusBadRequest, e.ERROR_WECHAT_CODE_FAIL, nil, "")
		return
	}

	if response.OpenId == "" {
		app.Response(c, http.StatusBadRequest, e.ERROR_WECHAT_REQUEST_FAIL, nil, "")
		return
	}

	data, responseCode := api.userService.CheckByCode(response.OpenId, response.SessionKey, response.UnionId)
	app.Response(c, http.StatusOK, responseCode, data, "")
}

func (api *WechatApi) WebCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	state := c.DefaultQuery("state", "")
	timestamp := com.StrTo(c.DefaultQuery("timestamp", "0")).MustInt64()
	// 验证合法性
	redirectUrl := setting.AppSetting.AppHost + "/wechat/web-callback?timestamp=" + strconv.FormatInt(timestamp, 10)
	sign := util.StringToMd5(redirectUrl)
	if strings.Compare(state, sign) != 0 {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil, "")
		return
	}
	nowTime := time.Now().Unix()
	expireTime := timestamp + (60 * 10)
	if expireTime < nowTime {
		app.Response(c, http.StatusBadRequest, e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, nil, "")
		return
	}
	if code == "" {
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil, "")
		return
	}
	response, err := api.wechat.GetWebAccessToken(code)
	if err != nil {
		app.Response(c, http.StatusBadRequest, e.ERROR_WECHAT_REQUEST_FAIL, nil, "")
		return
	}
	// 获取用户信息
	userInfo, _ := api.wechat.GetUserInfo(response.AccessToken, response.OpenId)

	data, responseCode := api.userService.WebScanLogin(userInfo)
	app.Response(c, http.StatusOK, responseCode, data, "")
}

func (api *WechatApi) GetWechatLoginUrl(c *gin.Context) {
	requestUrl := api.wechat.GetWebCode()
	c.Redirect(http.StatusMovedPermanently, requestUrl)
}
