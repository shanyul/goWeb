package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_LOGIN_PARAMS = 10000

	ERROR_EXIST_CAT       = 10001
	ERROR_NOT_EXIST_CAT   = 10002
	ERROR_NOT_EXIST_WORKS = 10003
	ERROR_EXIST_WORKS     = 10004

	ERROR_NOT_EXIST   = 10005
	ERROR_GET_FAIL    = 10006
	ERROR_SEARCH_FAIL = 10007
	ERROR_ADD_FAIL    = 10008
	ERROR_EDIT_FAIL   = 10009
	ERROR_DELETE_FAIL = 10010

	ERROR_ADD_WORKS_FAIL         = 10013
	ERROR_COUNT_WORKS_FAIL       = 10014
	ERROR_GET_WORKS_FAIL         = 10015
	ERROR_CHECK_EXIST_WORKS_FAIL = 10016
	ERROR_DELETE_WORKS_FAIL      = 10017
	ERROR_UPLOAD_FILE_NOT_INPUT  = 10018

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
	ERROR_TOKEN_NOT_EXIST          = 20005
	ERROR_USER_NICKNAME_EXIST      = 20006
	ERROR_CONFIRM_PASSWORD_NOT_EQ  = 20007
	ERROR_REGISTER_FAIL            = 20008
	ERROR_GENERATE_PASSWORD        = 20009

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_LOGIN_PARAMS:              "用户名或密码错误",
	ERROR_EXIST_CAT:                 "已存在该分类名称",
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
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_TOKEN_NOT_EXIST:           "缺少用户验证参数",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
	ERROR_ADD_WORKS_FAIL:            "新增作品失败",
	ERROR_COUNT_WORKS_FAIL:          "获取记录总数失败",
	ERROR_GET_WORKS_FAIL:            "获取作品失败",
	ERROR_CHECK_EXIST_WORKS_FAIL:    "作品不存在",
	ERROR_DELETE_WORKS_FAIL:         "作品删除失败",
	ERROR_USER_NICKNAME_EXIST:       "昵称已存在",
	ERROR_CONFIRM_PASSWORD_NOT_EQ:   "确认密码错误",
	ERROR_REGISTER_FAIL:             "注册失败",
	ERROR_GENERATE_PASSWORD:         "生成密码失败，请重试",
	ERROR_UPLOAD_FILE_NOT_INPUT:     "请选择上传文件",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
