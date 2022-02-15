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

// 获取多个评论
func GetTopics(c *gin.Context) {
	appG := app.Gin{C: c}
	var topicService service.Topic

	if worksId := c.Query("worksId"); worksId != "" {
		topicService.WorksId = com.StrTo(worksId).MustInt()
	}
	if userId := c.Query("userId"); userId != "" {
		topicService.UserId = com.StrTo(userId).MustInt()
	}

	topicService.PageNum = util.GetPage(c)
	topicService.PageSize = setting.AppSetting.PageSize

	total, err := topicService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_WORKS_FAIL, nil)
		return
	}

	topic, err := topicService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = topic
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// AddTopic 新增文章作品
func AddTopic(c *gin.Context) {
	var (
		appG         = app.Gin{C: c}
		form         request.AddTopicForm
		topicService service.Topic
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	if relationId := c.DefaultPostForm("relationId", "0"); relationId != "" {
		topicService.RelationId = com.StrTo(relationId).MustInt()
	}

	// 获取用户信息
	id := (c.MustGet("userId")).(int)
	userInfo := service.GetUserInfo(id)

	topicService.WorksId = form.WorksId
	topicService.Content = form.Content
	topicService.UserId = userInfo.UserId
	topicService.Username = userInfo.Username

	if err := topicService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// 删除评论
func DeleteTopic(c *gin.Context) {
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
	userId := (c.MustGet("userId")).(int)
	userInfo := service.GetUserInfo(userId)

	topicService := service.Topic{
		TopicId: id,
		UserId:  userInfo.UserId,
	}

	if err := topicService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
