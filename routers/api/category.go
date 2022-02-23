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

type CategoryApi struct {
	categoryService service.CategoryService
}

//获取多个作品
func (api *CategoryApi) GetCategory(c *gin.Context) {
	valid := validation.Validation{}

	name := c.DefaultQuery("name", "")
	parentId := -1
	if arg := c.Query("parentId"); arg != "" {
		parentId = com.StrTo(arg).MustInt()
		valid.Min(parentId, 1, "parentId")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil, "")
		return
	}

	categoryData := service.Category{}
	categoryData.CatName = name
	categoryData.ParentId = parentId
	categoryData.PageNum = util.GetPage(c)
	categoryData.PageSize = setting.AppSetting.PageSize

	total, err := api.categoryService.Count(&categoryData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil, "")
		return
	}

	category, err := api.categoryService.GetAll(&categoryData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil, "")
		return
	}

	data := make(map[string]interface{})
	data["lists"] = category
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data, "")
}

//新增文章作品
func (api *CategoryApi) AddCategory(c *gin.Context) {
	var (
		form request.AddCategoryForm
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	categoryData := service.Category{}
	categoryData.CatName = form.CatName
	categoryData.ParentId = form.ParentId

	if err := api.categoryService.Add(&categoryData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

//修改文章作品
func (api *CategoryApi) EditCategory(c *gin.Context) {
	var (
		form = request.EditCategoryForm{CatId: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	categoryData := service.Category{}
	categoryData.CatId = form.CatId
	categoryData.CatName = form.CatName
	categoryData.ParentId = form.ParentId

	exists, err := api.categoryService.ExistByID(form.CatId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST, nil, "")
		return
	}

	if err := api.categoryService.Edit(&categoryData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

//删除文章作品
func (api *CategoryApi) DeleteCategory(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil, "")
		return
	}

	exists, err := api.categoryService.ExistByID(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST, nil, "")
		return
	}

	err = api.categoryService.Delete(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}
