package util

import (
	"designer-api/pkg/setting"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type PageTool struct{}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	Start     int `json:"start"`
	Total     int `json:"total"`
	TotalPage int `json:"totalPage"`
}

// GetPage 获取初始页码
func (*PageTool) GetPage(c *gin.Context) Pagination {
	var pageInfo Pagination
	pageInfo.Page = com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("pageSize", "0")).MustInt()
	if pageSize == 0 {
		pageSize = setting.AppSetting.PageSize
	}
	pageInfo.PageSize = pageSize
	pageInfo.Start = (pageInfo.Page - 1) * pageInfo.PageSize

	return pageInfo
}
