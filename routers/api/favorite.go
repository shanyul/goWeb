package api

import (
	"designer-api/internal/request"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type FavoriteApi struct {
	favoriteService service.FavoriteService
}

// 获取多个关注
func (api *FavoriteApi) GetFavorite(c *gin.Context) {
	appG := app.Gin{C: c}
	var favoriteData service.Favorite

	if worksId := c.Query("worksId"); worksId != "" {
		favoriteData.WorksId = com.StrTo(worksId).MustInt()
	}
	if userId := c.Query("userId"); userId != "" {
		favoriteData.UserId = com.StrTo(userId).MustInt()
	}

	total, err := api.favoriteService.Count(&favoriteData)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	favorite, err := api.favoriteService.GetAll(&favoriteData)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = favorite
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// AddFavorite 添加关注
func (api *FavoriteApi) AddFavorite(c *gin.Context) {
	var (
		appG         = app.Gin{C: c}
		form         request.AddFavoriteForm
		favoriteData service.Favorite
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)

	favoriteData.WorksId = form.WorksId
	favoriteData.UserId = id

	if err := api.favoriteService.Add(&favoriteData); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 取消关注
func (api *FavoriteApi) DeleteFavorite(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	// 获取用户信息
	userId := (c.MustGet("userId")).(int)

	favoriteData := service.Favorite{}
	favoriteData.WorksId = id
	favoriteData.UserId = userId

	if err := api.favoriteService.Delete(&favoriteData); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
