package app

import (
	"designer-api/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

var MessageTmp = map[string]string{
	"Required":     "不能为空",
	"Min":          "最小值为 %d",
	"Max":          "最大值为 %d",
	"Range":        "范围必须是 %d 到 %d",
	"MinSize":      "最小值为 %d",
	"MaxSize":      "最大值为 %d",
	"Length":       "长度必须是 %d",
	"Alpha":        "只能输入字母",
	"Numeric":      "只能输入数字",
	"AlphaNumeric": "只能输入字母和数字",
	"Match":        "值必须为 %s",
	"NoMatch":      "值不能为 %s",
	"AlphaDash":    "只能输入字母、数字和下划线",
	"Email":        "请输入正确的邮箱",
	"IP":           "请输入正确的IP格式",
	"Base64":       "只能输入base64格式的字符",
	"Mobile":       "请输入正确的电话号码",
	"Tel":          "请输入正确的手机号码",
	"Phone":        "请输入正确的电话号码",
	"ZipCode":      "请填写正确的邮编",
}

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, form interface{}) (int, int, string) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS, ""
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR, ""
	}
	if !check {
		var errStr string
		for _, err := range valid.Errors {
			errStr += err.Message + ";"
		}
		return http.StatusBadRequest, e.INVALID_PARAMS, errStr
	}

	return http.StatusOK, e.SUCCESS, ""
}
