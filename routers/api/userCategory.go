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

type UserCategoryApi struct {
	userCategoryService service.UserCategoryService
	userService         service.UserService
}

//获取多个用户分类
func (api *UserCategoryApi) GetUserCategory(c *gin.Context) {
	var categoryData service.UserCategory
	if name := c.Query("name"); name != "" {
		categoryData.UcatName = name
	}

	if userId := c.Query("userId"); userId != "" {
		categoryData.UserId = com.StrTo(userId).MustInt()
	}
	if catId := c.Query("ucatId"); catId != "" {
		categoryData.UcatId = com.StrTo(catId).MustInt()
	}

	total, err := api.userCategoryService.Count(&categoryData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	works, err := api.userCategoryService.GetAll(&categoryData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = works
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

func (api *UserCategoryApi) GetOneCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	// 获取用户信息
	userId := (c.MustGet("userId")).(int)
	category, err := api.userCategoryService.Get(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, category)
}

func (api *UserCategoryApi) AddCategory(c *gin.Context) {
	var (
		form request.AddUserCategoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}
	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := api.userService.GetUserInfo(id)

	if exist, _ := api.userCategoryService.ExistByName(form.UcatName, id); exist {
		app.Response(c, httpCode, e.ERROR_NOT_EXIST_CAT, nil)
		return
	}

	categoryData := service.UserCategory{}
	categoryData.UcatName = form.UcatName
	categoryData.UserId = id
	categoryData.Username = userInfo.Username

	if err := api.userCategoryService.Add(&categoryData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func (api *UserCategoryApi) EditCategory(c *gin.Context) {
	var (
		form request.EditUserCategoryForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	if result, _ := api.userCategoryService.Get(form.UcatId, id); result.UcatId < 1 {
		app.Response(c, http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}

	categoryData := service.UserCategory{}
	categoryData.UcatId = form.UcatId
	categoryData.UcatName = form.UcatName

	if err := api.userCategoryService.Edit(&categoryData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func (api *UserCategoryApi) DeleteCategory(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	// 获取用户信息
	userId := (c.MustGet("userId")).(int)
	if result, _ := api.userCategoryService.Get(id, userId); result.UcatId < 1 {
		app.Response(c, http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}

	err := api.userCategoryService.Delete(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
