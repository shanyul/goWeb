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

//获取多个作品
func GetWorks(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	name := c.DefaultQuery("name", "")
	userId := com.StrTo(c.DefaultQuery("userId", "-1")).MustInt()

	catId := -1
	if arg := c.Query("catId"); arg != "" {
		catId = com.StrTo(arg).MustInt()
		valid.Min(catId, 1, "catId")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	worksService := service.Works{
		WorksName: name,
		CatId:     catId,
		UserId:    userId,
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}

	total, err := worksService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	works, err := worksService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = works
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetOneWorks(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	worksService := service.Works{
		WorksId: id,
	}

	exists, err := worksService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil)
		return
	}

	works, err := worksService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, works)
}

// AddWorks 新增文章作品
func AddWorks(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddWorksForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	worksService := service.Works{
		WorksName:        form.WorksName,
		UserId:           form.UserId,
		State:            form.State,
		CatId:            form.CatId,
		WorksLink:        form.WorksLink,
		WorksType:        form.WorksType,
		WorksDescription: form.WorksDescription,
		Remark:           form.Remark,
	}

	if err := worksService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_WORKS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 修改文章作品
func EditWorks(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.EditWorksForm{WorksId: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	worksService := service.Works{
		WorksId:          form.WorksId,
		WorksName:        form.WorksName,
		State:            form.State,
		CatId:            form.CatId,
		WorksLink:        form.WorksLink,
		WorksType:        form.WorksType,
		WorksDescription: form.WorksDescription,
		Remark:           form.Remark,
	}

	exists, err := worksService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil)
		return
	}

	if err := worksService.Edit(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_WORKS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 删除文章作品
func DeleteWorks(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	worksService := service.Works{
		WorksId: id,
	}

	exists, err := worksService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil)
		return
	}

	err = worksService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_WORKS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
