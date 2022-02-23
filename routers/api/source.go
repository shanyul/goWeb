package api

import (
	"designer-api/internal/request"
	"designer-api/internal/service"
	"designer-api/pkg/app"
	"designer-api/pkg/e"
	"designer-api/pkg/setting"
	"designer-api/pkg/util"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

type SourceApi struct {
	sourceService       service.UserSourceService
	userService         service.UserService
	userCategoryService service.UserCategoryService
}

//获取多个用户分类
func (api *SourceApi) GetSource(c *gin.Context) {
	var sourceData service.UserSource
	if ucatId := c.Query("ucatId"); ucatId != "" {
		sourceData.UcatId = com.StrTo(ucatId).MustInt()
	}

	if userId := c.Query("userId"); userId != "" {
		sourceData.UserId = com.StrTo(userId).MustInt()
	}
	if ucatName := c.Query("ucatName"); ucatName != "" {
		sourceData.UcatName = ucatName
	}
	if title := c.Query("title"); title != "" {
		sourceData.Title = title
	}

	total, err := api.sourceService.Count(&sourceData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil, "")
		return
	}

	sourceData.PageNum = util.GetPage(c)
	sourceData.PageSize = setting.AppSetting.PageSize

	sourceList, err := api.sourceService.GetAll(&sourceData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil, "")
		return
	}

	data := make(map[string]interface{})
	data["lists"] = sourceList
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data, "")
}

func (api *SourceApi) GetOneSource(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil, "")
		return
	}

	category, err := api.sourceService.Get(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, category, "")
}

func (api *SourceApi) AddSource(c *gin.Context) {
	var (
		form request.AddUserSourceForm
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := api.userService.GetUserInfo(id)

	category, err := api.userCategoryService.Get(form.UcatId, id)
	if category.UcatId == 0 || err != nil {
		app.Response(c, httpCode, e.ERROR_NOT_EXIST_CAT, nil, "")
		return
	}

	sourceData := service.UserSource{}
	sourceData.UserId = id
	sourceData.Username = userInfo.Username
	sourceData.UcatName = category.UcatName
	sourceData.UcatId = form.UcatId
	sourceData.Description = form.Description
	sourceData.Title = form.Title
	sourceData.Link = form.Link

	if err := api.sourceService.Add(&sourceData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

func (api *SourceApi) EditSource(c *gin.Context) {
	var (
		form request.EditUserSourceForm
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	category, err := api.userCategoryService.Get(form.UcatId, id)
	if category.UcatId == 0 || err != nil {
		app.Response(c, httpCode, e.ERROR_NOT_EXIST_CAT, nil, "")
		return
	}

	sourceData := service.UserSource{}
	sourceData.UcatName = category.UcatName
	sourceData.SourceId = form.SourceId
	sourceData.UcatId = form.UcatId
	sourceData.Description = form.Description
	sourceData.Title = form.Title
	sourceData.Link = form.Link

	if err := api.sourceService.Edit(&sourceData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

func (api *SourceApi) DeleteSource(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil, "")
		return
	}
	// 获取用户信息
	userId := (c.MustGet("userId")).(int)
	result, err := api.sourceService.Get(id)
	if result.UserId != userId {
		app.Response(c, http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil, "")
		return
	}

	err = api.sourceService.Delete(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}
