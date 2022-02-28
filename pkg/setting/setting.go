package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	AppName         string
	AppHost         string
	PageSize        int
	JwtSecret       string
	SignKey         string
	Environment     string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExt   []string
	RuntimeRootPath string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Log struct {
	RootPath   string
	SavePath   string
	SaveName   string
	FileExt    string
	TimeFormat string
}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
	Timeout  time.Duration
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Wechat struct {
	AppId        string
	AppSecret    string
	WebAppId     string
	WebAppSecret string
	CosKey       string
	CosSecret    string
}

var (
	cfg             *ini.File
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	LogSetting      = &Log{}
	DatabaseSetting = &Database{}
	RedisSetting    = &Redis{}
	WechatSetting   = &Wechat{}
)

// Setup 初始化配置
func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("log", LogSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("wechat", WechatSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	DatabaseSetting.Timeout = DatabaseSetting.Timeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

// 参数赋值
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
