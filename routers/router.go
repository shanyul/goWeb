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
	// 文件处理,上传操作需要登录
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.POST("/upload/image", middleware.JWT(), baseApi.UploadApi.UploadImage)
	r.POST("/upload", middleware.JWT(), baseApi.UploadApi.Upload)

	// 网页用户登录注册
	r.POST("/auth/login", baseApi.UserApi.Login)
	r.POST("/auth/register", baseApi.UserApi.Register)
	// 网页微信扫码登录
	r.GET("/wechat/web-login", baseApi.wechatApi.GetWechatLoginUrl)
	r.GET("/wechat/web-callback", baseApi.wechatApi.WebCallback)
	// 小程序登录
	r.GET("/wechat/login", baseApi.wechatApi.Login)
	// 验证码
	r.GET("/captcha/show/:image", baseApi.CaptchaApi.Show)
	r.GET("/captcha", baseApi.CaptchaApi.Get)
	// 用户操作
	r.GET("/auth/:id", middleware.Sign(), baseApi.UserApi.GetUserInfo)
	r.PUT("/auth/edit", middleware.JWT(), baseApi.UserApi.EditUser)
	r.PUT("/auth/change-password", middleware.JWT(), baseApi.UserApi.ChangePassword)
	r.GET("/refresh-token", middleware.JWT(), baseApi.UserApi.RefreshToken)

	apiHandle := r.Group("/api")
	{
		// ********* Sign 验证 **********
		apiHandle.GET("/works", middleware.Sign(), baseApi.WorksApi.GetWorks)
		apiHandle.GET("/works/:id", middleware.Sign(), baseApi.WorksApi.GetOneWorks)
		apiHandle.GET("/topic", middleware.Sign(), baseApi.TopicApi.GetTopics)
		apiHandle.GET("/tags", middleware.Sign(), baseApi.tagsApi.GetList)
		apiHandle.GET("/tags/:id", middleware.Sign(), baseApi.tagsApi.GetOne)
		apiHandle.GET("/favorite/add", middleware.Sign(), baseApi.FavoriteApi.AddFavorite)

		// ********* Jwt 验证 **********
		// 类别
		apiHandle.GET("/category", middleware.Sign(), baseApi.CategoryApi.GetCategory)
		apiHandle.POST("/category", middleware.JWT(), baseApi.CategoryApi.AddCategory)
		apiHandle.PUT("/category/:id", middleware.JWT(), baseApi.CategoryApi.EditCategory)
		apiHandle.DELETE("/category/:id", middleware.JWT(), baseApi.CategoryApi.DeleteCategory)
		// 作品
		apiHandle.POST("/works", middleware.JWT(), baseApi.WorksApi.AddWorks)
		apiHandle.PUT("/works", middleware.JWT(), baseApi.WorksApi.EditWorks)
		apiHandle.PUT("/works/recover/:id", middleware.JWT(), baseApi.WorksApi.Recover)
		apiHandle.DELETE("/works/:id", middleware.JWT(), baseApi.WorksApi.DeleteWorks)
		apiHandle.DELETE("/works/true/:id", middleware.JWT(), baseApi.WorksApi.Delete)
		// 评论
		apiHandle.POST("/topic", middleware.JWT(), baseApi.TopicApi.AddTopic)
		apiHandle.DELETE("/topic/:id", middleware.JWT(), baseApi.TopicApi.DeleteTopic)
		// 关注
		apiHandle.GET("/favorite", middleware.JWT(), baseApi.FavoriteApi.GetFavorite)
		apiHandle.DELETE("/favorite/:id", middleware.JWT(), baseApi.FavoriteApi.DeleteFavorite)
		// 用户类别
		apiHandle.GET("/user-category", middleware.JWT(), baseApi.userCategoryApi.GetUserCategory)
		apiHandle.GET("/user-category/:id", middleware.JWT(), baseApi.userCategoryApi.GetOneCategory)
		apiHandle.POST("/user-category", middleware.JWT(), baseApi.userCategoryApi.AddCategory)
		apiHandle.PUT("/user-category", middleware.JWT(), baseApi.userCategoryApi.EditCategory)
		apiHandle.DELETE("/user-category/:id", middleware.JWT(), baseApi.userCategoryApi.DeleteCategory)
		// 素材
		apiHandle.GET("/source", middleware.JWT(), baseApi.sourceApi.GetSource)
		apiHandle.GET("/source/:id", middleware.JWT(), baseApi.sourceApi.GetOneSource)
		apiHandle.POST("/source", middleware.JWT(), baseApi.sourceApi.AddSource)
		apiHandle.PUT("/source", middleware.JWT(), baseApi.sourceApi.EditSource)
		apiHandle.DELETE("/source/:id", middleware.JWT(), baseApi.sourceApi.DeleteSource)
		// 配置
		apiHandle.GET("/config", middleware.JWT(), baseApi.configApi.GetList)
		apiHandle.GET("/config/:key", middleware.JWT(), baseApi.configApi.GetOne)
		apiHandle.POST("/config", middleware.JWT(), baseApi.configApi.AddConfig)
		apiHandle.PUT("/config", middleware.JWT(), baseApi.configApi.EditConfig)
		apiHandle.DELETE("/config/:id", middleware.JWT(), baseApi.configApi.DeleteConfig)
		// 标签
		apiHandle.POST("/tags", middleware.JWT(), baseApi.tagsApi.AddTag)
		apiHandle.PUT("/tags", middleware.JWT(), baseApi.tagsApi.EditTag)
		apiHandle.DELETE("/tags/:id", middleware.JWT(), baseApi.tagsApi.Delete)
	}

	return r
}
