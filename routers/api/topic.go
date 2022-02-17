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

type TopicApi struct {
	topicService service.TopicService
	userService  service.UserService
}

// 获取多个评论
func (api *TopicApi) GetTopics(c *gin.Context) {
	var topicData service.Topic

	if worksId := c.Query("worksId"); worksId != "" {
		topicData.WorksId = com.StrTo(worksId).MustInt()
	}
	if userId := c.Query("userId"); userId != "" {
		topicData.UserId = com.StrTo(userId).MustInt()
	}

	total, err := api.topicService.Count(&topicData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}
	topicData.PageNum = util.GetPage(c)
	topicData.PageSize = setting.AppSetting.PageSize

	topic, err := api.topicService.GetAll(&topicData)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = topic
	data["total"] = total

	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

// AddTopic 新增文章作品
func (api *TopicApi) AddTopic(c *gin.Context) {
	var (
		form      request.AddTopicForm
		topicData service.Topic
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	if relationId := c.DefaultPostForm("relationId", "0"); relationId != "" {
		topicData.RelationId = com.StrTo(relationId).MustInt()
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := api.userService.GetUserInfo(id)

	topicData.WorksId = form.WorksId
	topicData.Content = form.Content
	topicData.RelationId = form.RelationId
	topicData.UserId = userInfo.UserId
	topicData.Username = userInfo.Nickname

	if err := api.topicService.Add(&topicData); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

// 删除评论
func (api *TopicApi) DeleteTopic(c *gin.Context) {
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

	if err := api.topicService.Delete(id, userId); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
