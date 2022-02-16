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
	userApi := api.UserApi{}
	r.POST("/auth/login", userApi.Login)
	r.POST("/auth/register", userApi.Register)
	r.PUT("/auth/edit", middleware.JWT(), userApi.EditUser)
	r.GET("/refresh-token", middleware.JWT(), userApi.RefreshToken)
	r.POST("/upload", middleware.JWT(), api.UploadImage)

	apiHandle := r.Group("/api")
	apiHandle.Use(middleware.JWT())
	{
		// 类别
		categoryApi := api.CategoryApi{}
		apiHandle.GET("/category", categoryApi.GetCategory)
		apiHandle.POST("/category", categoryApi.AddCategory)
		apiHandle.PUT("/category/:id", categoryApi.EditCategory)
		apiHandle.DELETE("/category/:id", categoryApi.DeleteCategory)
		// 作品
		worksApi := api.WorksApi{}
		apiHandle.GET("/works", worksApi.GetWorks)
		apiHandle.GET("/works/:id", worksApi.GetOneWorks)
		apiHandle.POST("/works", worksApi.AddWorks)
		apiHandle.PUT("/works/:id", worksApi.EditWorks)
		apiHandle.DELETE("/works/:id", worksApi.DeleteWorks)
		// 评论
		topicApi := api.TopicApi{}
		apiHandle.GET("/topic", topicApi.GetTopics)
		apiHandle.POST("/topic", topicApi.AddTopic)
		apiHandle.DELETE("/topic/:id", topicApi.DeleteTopic)
		// 关注
		favoriteApi := api.FavoriteApi{}
		apiHandle.GET("/favorite", favoriteApi.GetFavorite)
		apiHandle.POST("/favorite", favoriteApi.AddFavorite)
		apiHandle.DELETE("/favorite/:id", favoriteApi.DeleteFavorite)
	}

	return r
}
