package routers

import (
	"designer-api/middleware"
	"designer-api/pkg/setting"
	"designer-api/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiHandle := r.Group("/api")
	apiHandle.Use(middleware.JWT())
	{
		// 类别
		apiHandle.GET("/category", api.GetCategory)
		apiHandle.POST("/category", api.AddCategory)
		apiHandle.PUT("/category/:id", api.EditCategory)
		apiHandle.DELETE("/category/:id", api.DeleteCategory)
		// 作品
		apiHandle.GET("/works", api.GetWorks)
		apiHandle.GET("/works/:id", api.GetOneWorks)
		apiHandle.POST("/works", api.AddWorks)
		apiHandle.PUT("/works/:id", api.EditWorks)
		apiHandle.DELETE("/works/:id", api.DeleteWorks)
	}

	return r
}
