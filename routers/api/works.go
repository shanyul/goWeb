package api

import (
	"designer-api/models"
	"designer-api/pkg/e"
	"designer-api/pkg/setting"
	"designer-api/pkg/util"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

//获取多个作品
func GetWorks(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	valid := validation.Validation{}

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	maps["is_open"] = 1

	var catId int = -1
	if arg := c.Query("catId"); arg != "" {
		catId = com.StrTo(arg).MustInt()
		maps["cat_id"] = catId

		valid.Min(catId, 1, "cat_id").Message("类别ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetWorks(util.GetPage(c), setting.PageSize, maps)
		data["total"] = models.GetWorksTotal(maps)

	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetOneWorks(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistWorksById(id) {
			data = models.GetOneWorks(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_WORKS
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章作品
func AddWorks(c *gin.Context) {
	name := c.Query("name")
	userId := com.StrTo(c.Query("userId")).MustInt()
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	catId := com.StrTo(c.Query("catId")).MustInt()
	link := c.Query("link")
	workType := com.StrTo(c.DefaultQuery("workType", "0")).MustInt()
	description := c.Query("description")
	remark := com.StrTo(c.DefaultQuery("remark", "")).String()

	valid := validation.Validation{}
	valid.Min(catId, 1, "cat_id").Message("类别ID必须大于0")
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(userId, "userId").Message("userId不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(link, "link").Message("作品链接不能为空")
	valid.Required(description, "description").Message("作品描述不能为空")
	valid.Required(workType, "workType").Message("作品类型不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistWorksByName(name) {
			code = e.SUCCESS
			data := make(map[string]interface{})
			data["works_name"] = name
			data["user_id"] = userId
			data["state"] = state
			data["cat_id"] = catId
			data["works_link"] = link
			data["works_type"] = workType
			data["works_description"] = description
			data["remark"] = remark
			models.AddWorks(data)
		} else {
			code = e.ERROR_EXIST_WORKS
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章作品
func EditWorks(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	userId := com.StrTo(c.Query("userId")).MustInt()
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	catId := com.StrTo(c.Query("catId")).MustInt()
	link := c.Query("link")
	workType := com.StrTo(c.DefaultQuery("workType", "0")).MustInt()
	description := c.Query("description")
	remark := com.StrTo(c.DefaultQuery("remark", "")).String()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("修改项ID不能为空")
	valid.Min(catId, 1, "cat_id").Message("类别ID必须大于0")
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(userId, "userId").Message("userId不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(link, "link").Message("作品链接不能为空")
	valid.Required(description, "description").Message("作品描述不能为空")
	valid.Required(workType, "workType").Message("作品类型不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if models.ExistWorksById(id) {
			code = e.SUCCESS
			data := make(map[string]interface{})
			data["works_name"] = name
			data["user_id"] = userId
			data["state"] = state
			data["cat_id"] = catId
			data["works_link"] = link
			data["works_type"] = workType
			data["works_description"] = description
			data["remark"] = remark
			models.EditWorks(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_WORKS
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//删除文章作品
func DeleteWorks(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistWorksById(id) {
			models.DeleteWorks(id)
		} else {
			code = e.ERROR_NOT_EXIST_WORKS
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
