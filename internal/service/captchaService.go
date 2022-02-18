package service

import (
	"designer-api/pkg/setting"
	"github.com/dchest/captcha"
)

type CaptchaService struct{}

type Captcha struct {
	CaptchaId string `json:"captchaId"`
	ImageUrl  string `json:"imageUrl"`
	WavUrl    string `json:"wavUrl"`
}

func (service *CaptchaService) Get() Captcha {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)

	result := Captcha{}
	if captchaId != "" {
		result.CaptchaId = captchaId
		result.ImageUrl = setting.AppSetting.AppHost + "/captcha/show/" + captchaId + ".png"
		result.WavUrl = setting.AppSetting.AppHost + "/captcha/show/" + captchaId + ".wav"
	}

	return result
}

func (service *CaptchaService) Verify(captchaId string, code string) bool {
	return captcha.VerifyString(captchaId, code)
}
