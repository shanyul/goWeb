package routers

import (
	"designer-api/middleware"
	"designer-api/pkg/upload"
	"designer-api/routers/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/auth/login", api.Login)
	r.POST("/auth/register", api.Register)
	r.PUT("/auth/edit", middleware.JWT(), api.EditUser)
	r.GET("/refresh-token", middleware.JWT(), api.RefreshToken)
	r.POST("/upload", middleware.JWT(), api.UploadImage)

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
		// 评论
		apiHandle.GET("/topic", api.GetTopics)
		apiHandle.POST("/topic", api.AddTopic)
		apiHandle.DELETE("/topic/:id", api.DeleteTopic)
		// 关注
		apiHandle.GET("/favorite", api.GetFavorite)
		apiHandle.POST("/favorite", api.AddFavorite)
		apiHandle.DELETE("/favorite/:id", api.DeleteFavorite)
	}

	return r
}
