package request

// AddTagsForm 标签
type AddTagsForm struct {
	TagName string `form:"tagName" valid:"Required"`
}

// EditTagsForm 标签
type EditTagsForm struct {
	TagId   int    `form:"tagId" valid:"Required;Min(1)"`
	TagName string `form:"tagName" valid:"Required"`
}
