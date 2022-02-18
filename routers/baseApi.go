package routers

import "designer-api/routers/api"

type ApiCommon struct {
	UserApi         api.UserApi
	CategoryApi     api.CategoryApi
	FavoriteApi     api.FavoriteApi
	TopicApi        api.TopicApi
	WorksApi        api.WorksApi
	UploadApi       api.UploadApi
	CaptchaApi      api.CaptchaApi
	userCategoryApi api.UserCategoryApi
	sourceApi       api.SourceApi
}
