package request

// ConfigForm 配置
type ConfigForm struct {
	Key   string `form:"key" valid:"Required"`
	Value string `form:"value" valid:"Required"`
}
