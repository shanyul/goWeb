package api

import (
	"designer-api/models"
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
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	code := e.SUCCESS

	data["lists"] = models.GetCategory(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetCategoryTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章作品
func AddCategory(c *gin.Context) {
	name := c.Query("name")
	parentId := com.StrTo(c.DefaultQuery("parentId", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistCategoryByName(name) {
			code = e.SUCCESS
			models.AddCategory(name, parentId)
		} else {
			code = e.ERROR_EXIST_CAT
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章作品
func EditCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")

	valid := validation.Validation{}

	var parentId int = 0
	if arg := c.Query("parentId"); arg != "" {
		parentId = com.StrTo(arg).MustInt()
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCategoryById(id) {
			data := make(map[string]interface{})
			data["cat_name"] = name
			if parentId != 0 {
				data["parent_id"] = parentId
			}

			models.EditCategory(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_CAT
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章作品
func DeleteCategory(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistCategoryById(id) {
			models.DeleteCategory(id)
		} else {
			code = e.ERROR_NOT_EXIST_CAT
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
