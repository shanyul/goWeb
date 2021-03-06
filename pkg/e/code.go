package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_LOGIN_PARAMS    = 10000
	ERROR_LOGIN_FAIL      = 10001
	ERROR_NOT_EXIST_CAT   = 10002
	ERROR_NOT_EXIST_WORKS = 10003
	ERROR_EXIST_WORKS     = 10004

	ERROR_NOT_EXIST          = 10005
	ERROR_GET_FAIL           = 10006
	ERROR_SEARCH_FAIL        = 10007
	ERROR_ADD_FAIL           = 10008
	ERROR_EDIT_FAIL          = 10009
	ERROR_DELETE_FAIL        = 10010
	ERROR_GET_CAPTCHA_FAIL   = 10011
	ERROR_CHECK_CAPTCHA_FAIL = 10012

	ERROR_ADD_WORKS_FAIL         = 10013
	ERROR_COUNT_WORKS_FAIL       = 10014
	ERROR_GET_WORKS_FAIL         = 10015
	ERROR_CHECK_EXIST_WORKS_FAIL = 10016
	ERROR_DELETE_WORKS_FAIL      = 10017
	ERROR_UPLOAD_FILE_NOT_INPUT  = 10018
	ERROR_EXIST_FAIL             = 10019
	ERROR_OLD_PASSWORD_FAIL      = 10020
	ERROR_PASSWORD_CHANGE_FAIL   = 10021

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_TOKEN_NOT_EXIST          = 20005
	ERROR_USER_NICKNAME_EXIST      = 20006
	ERROR_CONFIRM_PASSWORD_NOT_EQ  = 20007
	ERROR_REGISTER_FAIL            = 20008
	ERROR_GENERATE_PASSWORD        = 20009
	ERROR_AUTH_CHECK_SIGN_FAIL     = 20010
	ERROR_AUTH_CHECK_SIGN_EXPIRE   = 20011

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
	ERROR_UPLOAD_CHECK_SIZE_FAIL    = 30004
	ERROR_UPLOAD_FAIL               = 30005

	ERROR_WECHAT_REQUEST_FAIL = 40000
	ERROR_WECHAT_CODE_FAIL    = 40001
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_LOGIN_PARAMS:              "用户名或密码错误",
	ERROR_LOGIN_FAIL:                "登录失败，请重试",
	ERROR_NOT_EXIST_CAT:             "该分类不存在",
	ERROR_NOT_EXIST_WORKS:           "该作品不存在",
	ERROR_NOT_EXIST:                 "数据不存在",
	ERROR_GET_FAIL:                  "获取失败",
	ERROR_SEARCH_FAIL:               "查询失败",
	ERROR_ADD_FAIL:                  "添加失败",
	ERROR_EDIT_FAIL:                 "编辑失败",
	ERROR_DELETE_FAIL:               "删除失败",
	ERROR_EXIST_WORKS:               "该作品已存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_SIGN_FAIL:      "签名验证失败",
	ERROR_AUTH_CHECK_SIGN_EXPIRE:    "签名验已过期",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_TOKEN_NOT_EXIST:           "请登录后操作",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	ERROR_ADD_WORKS_FAIL:            "新增作品失败",
	ERROR_COUNT_WORKS_FAIL:          "获取记录总数失败",
	ERROR_GET_WORKS_FAIL:            "获取作品失败",
	ERROR_CHECK_EXIST_WORKS_FAIL:    "作品不存在",
	ERROR_DELETE_WORKS_FAIL:         "作品删除失败",
	ERROR_USER_NICKNAME_EXIST:       "名称已存在",
	ERROR_CONFIRM_PASSWORD_NOT_EQ:   "确认密码错误",
	ERROR_REGISTER_FAIL:             "注册失败",
	ERROR_GENERATE_PASSWORD:         "生成密码失败，请重试",
	ERROR_UPLOAD_FILE_NOT_INPUT:     "请选择上传文件",
	ERROR_GET_CAPTCHA_FAIL:          "获取验证码失败",
	ERROR_CHECK_CAPTCHA_FAIL:        "验证码过期或错误",
	ERROR_EXIST_FAIL:                "数据已存在",
	ERROR_WECHAT_REQUEST_FAIL:       "请求微信接口失败",
	ERROR_WECHAT_CODE_FAIL:          "code无效",
	ERROR_OLD_PASSWORD_FAIL:         "旧密码错误",
	ERROR_PASSWORD_CHANGE_FAIL:      "密码修改失败",
	ERROR_UPLOAD_CHECK_SIZE_FAIL:    "请上传小于 20M 的文件",
	ERROR_UPLOAD_FAIL:               "文件上传失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
