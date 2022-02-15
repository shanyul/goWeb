package request

// EditUserForm 用户表单
type EditUserForm struct {
	Username string `form:"username" valid:"Required"`
	Nickname string `form:"nickname" valid:"Required"`
	Avatar   string `form:"avatar"`
	BgImage  string `form:"bgImage"`
	Province string `form:"province"`
	City     string `form:"city"`
	Distinct string `form:"distinct"`
	Remark   string `form:"remark"`
}
