package request

// AddUserSourceForm 分类表单
type AddUserSourceForm struct {
	UcatId      int    `form:"ucatId" valid:"Required;Min(1)"`
	UcatName    string `form:"ucatName" valid:"Required"`
	Description string `form:"description" valid:"Required"`
	Title       string `form:"title" valid:"Required"`
	Link        string `form:"link" valid:"Required"`
}

// EditUserSourceForm 分类表单
type EditUserSourceForm struct {
	SourceId    int    `form:"sourceId" valid:"Required;Min(1)"`
	UcatId      int    `form:"ucatId" valid:"Required;Min(1)"`
	UcatName    string `form:"ucatName" valid:"Required"`
	Description string `form:"description" valid:"Required"`
	Title       string `form:"title" valid:"Required"`
	Link        string `form:"link" valid:"Required"`
}
