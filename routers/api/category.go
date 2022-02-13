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
func GetCategory(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	name := c.DefaultQuery("name", "")
	parentId := -1
	if arg := c.Query("parentId"); arg != "" {
		parentId = com.StrTo(arg).MustInt()
		valid.Min(parentId, 1, "parentId")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	categoryService := service.Category{
		CatName:  name,
		ParentId: parentId,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := categoryService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	category, err := categoryService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = category
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//新增文章作品
func AddCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form request.AddCategoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := service.Category{
		CatName:  form.CatName,
		ParentId: form.ParentId,
	}

	if err := categoryService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//修改文章作品
func EditCategory(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = request.EditCategoryForm{CatId: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	categoryService := service.Category{
		CatId:    form.CatId,
		CatName:  form.CatName,
		ParentId: form.ParentId,
	}

	exists, err := categoryService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, nil)
		return
	}

	if err := categoryService.Edit(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//删除文章作品
func DeleteCategory(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	categoryService := service.Category{
		CatId: id,
	}

	exists, err := categoryService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, nil)
		return
	}

	err = categoryService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
