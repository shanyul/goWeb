package request

// EditUserForm 用户表单
type EditUserForm struct {
	Username     string `form:"username" valid:"Required"`
	Nickname     string `form:"nickname" valid:"Required"`
	Avatar       string `form:"avatar"`
	BgImage      string `form:"bgImage"`
	Province     string `form:"province"`
	City         string `form:"city"`
	Distinct     string `form:"distinct"`
	Address      string `form:"address"`
	Remark       string `form:"remark"`
	Profession   string `form:"profession"`
	Charge       string `form:"charge"`
	Introduction string `form:"introduction"`
}

type LoginUserForm struct {
	Username  string `form:"username" valid:"Required; AlphaDash; MaxSize(20)"`
	Password  string `form:"password" valid:"Required; MaxSize(20)"`
	Captcha   string `form:"captcha" valid:"Required"`
	CaptchaId string `form:"captchaId" valid:"Required"`
}

type RegisterUserForm struct {
	Username        string `form:"username" valid:"Required; AlphaDash; MaxSize(20)"`
	Nickname        string `form:"nickname" valid:"Required; MaxSize(20)"`
	Password        string `form:"password" valid:"Required; MaxSize(20)"`
	ConfirmPassword string `form:"confirmPassword" valid:"Required; MaxSize(20)"`
	Captcha         string `form:"captcha" valid:"Required"`
	CaptchaId       string `form:"captchaId" valid:"Required"`
}

type ChangePasswordForm struct {
	OldPassword     string `form:"oldPassword" valid:"Required; MaxSize(20)"`
	Password        string `form:"password" valid:"Required; MaxSize(20)"`
	ConfirmPassword string `form:"confirmPassword" valid:"Required; MaxSize(20)"`
}
