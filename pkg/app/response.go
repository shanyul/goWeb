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
func Response(ctx *gin.Context, httpCode, errCode int, data interface{}, msg string) {
	if msg == "" {
		msg = e.GetMsg(errCode)
	}
	ctx.JSON(httpCode, ResponseForm{
		Code: errCode,
		Msg:  msg,
		Data: data,
	})
	return
}
