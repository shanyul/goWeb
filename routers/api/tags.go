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

type TagsApi struct {
	tagsService service.TagService
	userService service.UserService
}

var tagOrderMap = map[string]string{
	"high": "count desc",
	"low":  "count asc",
}

// 获取标签
func (api *TagsApi) GetList(c *gin.Context) {
	var data service.Tags
	if name := c.Query("tagName"); name != "" {
		data.TagName = name
	}

	if isDelete := c.Query("deleteType"); isDelete != "" {
		data.Delete = com.StrTo(isDelete).MustInt()
	}

	if tagId := c.Query("tagId"); tagId != "" {
		data.TagId = com.StrTo(tagId).MustInt()
	}
	data.UserId = (c.MustGet("userId")).(int)

	orderString := ""
	if orderBy := c.Query("orderBy"); orderBy != "" {
		value, ok := tagOrderMap[orderBy]
		if ok {
			orderString = value
		}
	}
	data.OrderBy = orderString

	tags, err := api.tagsService.GetAll(&data)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil, "")
		return
	}

	response := make(map[string]interface{})
	response["lists"] = tags

	app.Response(c, http.StatusOK, e.SUCCESS, response, "")
}

func (api *TagsApi) GetOne(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil, "")
		return
	}

	// 获取用户信息
	tag, err := api.tagsService.Get(id)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, tag, "")
}

func (api *TagsApi) AddTag(c *gin.Context) {
	var form request.AddTagsForm

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	exist, _ := api.tagsService.ExistByName(form.TagName, id)
	if exist {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil, "")
		return
	}

	userInfo := api.userService.GetUserInfo(id)

	saveData := service.Tags{}
	saveData.TagName = form.TagName
	saveData.UserId = userInfo.UserId
	saveData.Username = userInfo.Nickname

	if err := api.tagsService.Add(&saveData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

// 修改
func (api *TagsApi) EditTag(c *gin.Context) {
	var form request.EditTagsForm

	httpCode, errCode, msg := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil, msg)
		return
	}
	userId := (c.MustGet("userId")).(int)

	saveData := service.Tags{}
	saveData.TagId = form.TagId
	saveData.TagName = form.TagName
	saveData.UserId = userId

	if err := api.tagsService.Edit(&saveData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}

// 删除
func (api *TagsApi) Delete(c *gin.Context) {
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

	err := api.tagsService.Delete(id, userId)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil, "")
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil, "")
}
