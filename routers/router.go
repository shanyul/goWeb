package routers

import (
	"designer-api/middleware"
	"designer-api/pkg/upload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 获取 API
	baseApi := ApiCommon{}
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/auth/login", baseApi.UserApi.Login)
	r.POST("/auth/register", baseApi.UserApi.Register)
	r.GET("/captcha", baseApi.CaptchaApi.Get)
	r.GET("/captcha/show/:image", baseApi.CaptchaApi.Show)
	r.GET("/auth/:id", middleware.JWT(), baseApi.UserApi.GetUserInfo)
	r.PUT("/auth/edit", middleware.JWT(), baseApi.UserApi.EditUser)
	r.GET("/refresh-token", middleware.JWT(), baseApi.UserApi.RefreshToken)
	r.POST("/upload", middleware.JWT(), baseApi.UploadApi.UploadImage)

	apiHandle := r.Group("/api")
	apiHandle.Use(middleware.JWT())
	{
		// 类别
		apiHandle.GET("/category", baseApi.CategoryApi.GetCategory)
		apiHandle.POST("/category", baseApi.CategoryApi.AddCategory)
		apiHandle.PUT("/category/:id", baseApi.CategoryApi.EditCategory)
		apiHandle.DELETE("/category/:id", baseApi.CategoryApi.DeleteCategory)
		// 作品
		apiHandle.GET("/works", baseApi.WorksApi.GetWorks)
		apiHandle.GET("/works/:id", baseApi.WorksApi.GetOneWorks)
		apiHandle.POST("/works", baseApi.WorksApi.AddWorks)
		apiHandle.PUT("/works", baseApi.WorksApi.EditWorks)
		apiHandle.DELETE("/works/:id", baseApi.WorksApi.DeleteWorks)
		// 评论
		apiHandle.GET("/topic", baseApi.TopicApi.GetTopics)
		apiHandle.POST("/topic", baseApi.TopicApi.AddTopic)
		apiHandle.DELETE("/topic/:id", baseApi.TopicApi.DeleteTopic)
		// 关注
		apiHandle.GET("/favorite", baseApi.FavoriteApi.GetFavorite)
		apiHandle.POST("/favorite", baseApi.FavoriteApi.AddFavorite)
		apiHandle.DELETE("/favorite/:id", baseApi.FavoriteApi.DeleteFavorite)
		// 用户类别
		apiHandle.GET("/user-category", baseApi.userCategoryApi.GetUserCategory)
		apiHandle.GET("/user-category/:id", baseApi.userCategoryApi.GetOneCategory)
		apiHandle.POST("/user-category", baseApi.userCategoryApi.AddCategory)
		apiHandle.PUT("/user-category", baseApi.userCategoryApi.EditCategory)
		apiHandle.DELETE("/user-category/:id", baseApi.userCategoryApi.DeleteCategory)
		// 素材
		apiHandle.GET("/source", baseApi.sourceApi.GetSource)
		apiHandle.GET("/source/:id", baseApi.sourceApi.GetOneSource)
		apiHandle.POST("/source", baseApi.sourceApi.AddSource)
		apiHandle.PUT("/source", baseApi.sourceApi.EditSource)
		apiHandle.DELETE("/source/:id", baseApi.sourceApi.DeleteSource)
		// 配置
		apiHandle.GET("/config", baseApi.configApi.GetList)
		apiHandle.GET("/config/:key", baseApi.configApi.GetOne)
		apiHandle.POST("/config", baseApi.configApi.AddConfig)
		apiHandle.PUT("/config", baseApi.configApi.EditConfig)
		apiHandle.DELETE("/config/:id", baseApi.configApi.DeleteConfig)
	}

	return r
}
