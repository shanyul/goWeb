package util

import (
	"crypto/md5"
	"designer-api/pkg/app"
	"designer-api/pkg/setting"
	"encoding/hex"
	"github.com/astaxie/beego/validation"
)

// Setup Initialize the util
func Setup() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
	// set valid
	validation.SetDefaultMessage(app.MessageTmp)
}

func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
