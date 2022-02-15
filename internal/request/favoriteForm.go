package request

// AddFavoriteForm 评论表单
type AddFavoriteForm struct {
	WorksId int `form:"worksId" valid:"Required;Min(1)"`
}
