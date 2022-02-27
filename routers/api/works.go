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

type WorksApi struct {
	worksService service.WorksService
	viewService  service.ViewService
	userService  service.UserService
}

var orderMap = map[string]string{
	"new":      "works_id desc",
	"favorite": "favorite_num desc",
	"topic":    "topic_num desc",
	"view":     "view_num desc",
}

//获取多个作品
func (api *WorksApi) GetWorks(c *gin.Context) {
	var (
		orderString string
		worksData   service.Works
	)
	if name := c.Query("name"); name != "" {
		worksData.WorksName = name
	}

	if designer := c.Query("designer"); designer != "" {
		worksData.Username = designer
	}
	if isDelete := c.Query("delete"); isDelete != "" {
		worksData.Delete = com.StrTo(isDelete).MustInt()
	}
	if isOpen := c.Query("isOpen"); isOpen != "" {
		worksData.IsOpen = com.StrTo(isOpen).MustInt()
	} else {
		worksData.IsOpen = 1
	}
	orderString = ""
	if orderBy := c.Query("orderBy"); orderBy != "" {
		value, ok := orderMap[orderBy]
		if ok {
			orderString = value
		}
	}

	total, err := api.worksService.Count(&worksData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil, "")
		return
	}

	worksData.OrderBy = orderString
	worksData.PageNum = util.GetPage(c)
	worksData.PageSize = setting.AppSetting.PageSize

	works, err := api.worksService.GetAll(&worksData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil, "")
		return
	}

	data := make(map[string]interface{})
	data["lists"] = works
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data, "")
}

func (api *WorksApi) GetOneWorks(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil, "")
		return
	}

	exists, err := api.worksService.ExistByID(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil, "")
		return
	}

	// 获取用户信息
	//userId := (c.MustGet("userId")).(int)
	userId := 0
	works, err := api.worksService.Get(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_WORKS_FAIL, nil, "")
		return
	}

	// 更新查看数
	viewData := service.Viewer{}
	viewData.UserId = userId
	viewData.WorksId = id

	_ = api.viewService.Add(&viewData)

	app.Response(c, http.StatusOK, e.SUCCESS, works, "")
}

// AddWorks 新增文章作品
func (api *WorksApi) AddWorks(c *gin.Context) {
	var form request.AddWorksForm
	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := api.userService.GetUserInfo(id)

	worksData := service.Works{}
	worksData.WorksName = form.WorksName
	worksData.UserId = userInfo.UserId
	worksData.Username = userInfo.Nickname
	worksData.State = form.State
	worksData.WorksLink = form.WorksLink
	worksData.WorksType = form.WorksType
	worksData.IsOpen = form.IsOpen
	worksData.WorksDescription = form.WorksDescription
	worksData.Remark = form.Remark
	// 标签
	worksData.TagId = form.TagId
	worksData.TagName = form.TagName

	if err := api.worksService.Add(&worksData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_WORKS_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

// 修改文章作品
func (api *WorksApi) EditWorks(c *gin.Context) {
	var (
		form request.EditWorksForm
	)

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	worksData := service.Works{}
	worksData.WorksId = form.WorksId
	worksData.WorksName = form.WorksName
	worksData.State = form.State
	worksData.IsOpen = form.IsOpen
	worksData.WorksLink = form.WorksLink
	worksData.WorksType = form.WorksType
	worksData.WorksDescription = form.WorksDescription
	worksData.Remark = form.Remark
	// 标签
	worksData.TagId = form.TagId
	worksData.TagName = form.TagName

	exists, err := api.worksService.ExistByID(form.WorksId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil, "")
		return
	}

	if err := api.worksService.Edit(&worksData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_WORKS_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

// 删除文章作品
func (api *WorksApi) DeleteWorks(c *gin.Context) {
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
	exists, err := api.worksService.ExistByID(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil, "")
		return
	}

	err = api.worksService.Delete(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_WORKS_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

func (api *WorksApi) Delete(c *gin.Context) {
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
	exists, err := api.worksService.ExistByID(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_WORKS_FAIL, nil, "")
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_WORKS, nil, "")
		return
	}

	err = api.worksService.TrueDelete(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_WORKS_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}
