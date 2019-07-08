package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	AppFile *ini.File
	Cfg *ini.File
	ENV string


	RunMode string


)
var RedisSetting = & RedisConf{}

type RedisConf struct {
	RedisAddress string
	RedisPassword string
	RedisDB int
}

type AppConf struct {
	PageSize int
	JwtSecret string
}
var AppSetting = & AppConf{}

type ServerConf struct {
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = & ServerConf{}

func init()  {
	var err error
	AppFile, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}
	ENV = AppFile.Section("").Key("ENV").MustString("dev")

	Cfg, err = ini.Load("conf/"+ENV+".ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/"+ENV+".ini': %v", err)
	}


	LoadBase()
	LoadServer()
	LoadApp()
	LoadRedis()
}


func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	err := Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}
}

func LoadApp() {
	err := Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}
}

func LoadRedis()  {
	err := Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}
}