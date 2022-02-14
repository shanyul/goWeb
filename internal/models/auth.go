package models

type User struct {
	UserID          int    `gorm:"primary_key" column:"user_id" json:"userId"`
	Username        string `column:"username" json:"username"`
	Password        string `column:"password" json:"password"`
	Nickname        string `column:"nickname" json:"nickname"`
	Avatar          string `column:"avatar" json:"avatar"`
	BgImage         string `column:"bg_image" json:"bgImage"`
	Phone           string `column:"phone" json:"phone"`
	Email           string `column:"email" json:"email"`
	State           string `column:"state" json:"state"`
	Province        string `column:"province" json:"province"`
	City            string `column:"city" json:"city"`
	Distinct        string `column:"distinct" json:"distinct"`
	Address         string `column:"address" json:"address"`
	Remark          string `column:"remark" json:"remark"`
	WechatOpenid    string `column:"wechat_openid" json:"wechatOpenid"`
	CreateTime      string `column:"create_time" json:"createTime"`
	UpdateTime      string `column:"update_time" json:"updateTime"`
	DeleteTimestamp int    `column:"delete_timestamp" json:"deleteTimestamp"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

// CheckAuth 验证用户
func CheckAuth(username, password string) (*User, bool) {
	var auth User
	err := dbHandle.Select(
		"user_id", "username", "nickname", "avatar", "bg_image", "phone", "email", "state", "province", "city", "distinct", "address", "create_time",
	).Where(User{Username: username, Password: password}).First(&auth).Error
	if err != nil {
		return nil, false
	}

	return &auth, auth.UserID > 0
}
