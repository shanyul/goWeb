package models

type User struct {
	UserID          int    `gorm:"primary_key" json:"user_id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	BgImage         string `json:"bg_image"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	State           string `json:"state"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Distinct        string `json:"distinct"`
	Address         string `json:"address"`
	Remark          string `json:"remark"`
	WechatOpenid    string `json:"wechat_openid"`
	CreateTime      string `json:"create_time"`
	UpdateTime      string `json:"update_time"`
	DeleteTimestamp int    `json:"delete_timestamp"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

// CheckAuth 验证用户
func CheckAuth(username, password string) bool {
	var auth User
	dbHandle.Select("user_id").Where(User{Username: username, Password: password}).First(&auth)

	return auth.UserID > 0
}
