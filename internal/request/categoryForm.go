package request

// AddCategoryForm 分类表单
type AddCategoryForm struct {
	CatName  string `form:"catName" valid:"Required"`
	ParentId int    `form:"parentId" valid:"Required;Min(1)"`
}

// EditCategoryForm 分类表单
type EditCategoryForm struct {
	CatId    int    `form:"catId" valid:"Required"`
	CatName  string `form:"catName" valid:"Required"`
	ParentId int    `form:"parentId" valid:"Required;Min(1)"`
}
