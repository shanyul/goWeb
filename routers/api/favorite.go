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

// 获取多个关注
func GetFavorite(c *gin.Context) {
	appG := app.Gin{C: c}
	var favoriteService service.Favorite

	if worksId := c.Query("worksId"); worksId != "" {
		favoriteService.WorksId = com.StrTo(worksId).MustInt()
	}
	if userId := c.Query("userId"); userId != "" {
		favoriteService.UserId = com.StrTo(userId).MustInt()
	}

	total, err := favoriteService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	favorite, err := favoriteService.GetAll()
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
func AddFavorite(c *gin.Context) {
	var (
		appG            = app.Gin{C: c}
		form            request.AddFavoriteForm
		favoriteService service.Favorite
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := service.GetUserInfo(id)

	favoriteService.WorksId = form.WorksId
	favoriteService.UserId = userInfo.UserId

	if err := favoriteService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 取消关注
func DeleteFavorite(c *gin.Context) {
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
	userid := (c.MustGet("userId")).(int)
	userInfo := service.GetUserInfo(userid)

	favoriteService := service.Favorite{
		WorksId: id,
		UserId:  userInfo.UserId,
	}

	if err := favoriteService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
