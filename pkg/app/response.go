package app

import (
	"designer-api/pkg/e"
	"github.com/gin-gonic/gin"
)

type ResponseForm struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func Response(ctx *gin.Context, httpCode, errCode int, data interface{}) {
	ctx.JSON(httpCode, ResponseForm{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
	return
}
