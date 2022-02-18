package request

// AddUserCategoryForm 分类表单
type AddUserCategoryForm struct {
	UcatName string `form:"ucatName" valid:"Required"`
}

// EditUserCategoryForm 分类表单
type EditUserCategoryForm struct {
	UcatId    int    `form:"ucatId" valid:"Required;Min(1)"`
	UcatName string `form:"ucatName" valid:"Required"`
}
