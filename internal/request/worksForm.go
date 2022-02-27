package request

// AddWorksForm 作品表单
type AddWorksForm struct {
	WorksName        string `form:"worksName" valid:"Required"`
	State            int    `form:"state" valid:"Min(0)"`
	IsOpen           int    `form:"isOpen" valid:"Range(0,1)"`
	WorksLink        string `form:"worksLink" valid:"Required;MaxSize(255)"`
	WorksType        int    `form:"worksType" valid:"Min(0)"`
	TagId            string `form:"tagId"`
	TagName          string `form:"tagName"`
	WorksDescription string `form:"worksDescription" valid:"Required"`
	Remark           string `form:"remark" valid:"MaxSize(255)"`
}

// EditWorksForm 作品表单
type EditWorksForm struct {
	WorksId          int    `form:"worksId" valid:"Required"`
	WorksName        string `form:"worksName" valid:"Required"`
	State            int    `form:"state" valid:"Min(0)"`
	IsOpen           int    `form:"isOpen" valid:"Range(0,1)"`
	TagId            string `form:"tagId"`
	TagName          string `form:"tagName"`
	WorksLink        string `form:"worksLink" valid:"Required;MaxSize(255)"`
	WorksType        int    `form:"worksType" valid:"Required"`
	WorksDescription string `form:"worksDescription" valid:"Required"`
	Remark           string `form:"remark" valid:"MaxSize(255)"`
}
