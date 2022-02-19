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

type ConfigApi struct {
	configService service.ConfigService
}

//获取多个用户分类
func (api *ConfigApi) GetList(c *gin.Context) {
	var configData service.Config
	if key := c.Query("key"); key != "" {
		configData.Key = key
	}

	total, err := api.configService.Count(&configData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	works, err := api.configService.GetAll(&configData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = works
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

func (api *ConfigApi) GetOne(c *gin.Context) {
	key := c.Param("key")
	// 获取用户信息
	config, err := api.configService.Get(key)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, config)
}

func (api *ConfigApi) AddConfig(c *gin.Context) {
	var (
		form request.ConfigForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}
	if exist, _ := api.configService.ExistByName(form.Key); exist {
		app.Response(c, httpCode, e.ERROR_NOT_EXIST_CAT, nil)
		return
	}

	data := service.Config{}
	data.Key = form.Key
	data.Value = form.Value

	if err := api.configService.Add(&data); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func (api *ConfigApi) EditConfig(c *gin.Context) {
	var (
		form request.ConfigForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	data := service.Config{}
	data.Key = form.Key
	data.Value = form.Value

	if err := api.configService.Edit(&data); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func (api *ConfigApi) DeleteConfig(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	err := api.configService.Delete(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
