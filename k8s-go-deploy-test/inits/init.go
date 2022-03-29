package inits

import (
	"flag"
	"github.com/kataras/iris/v12"
	"go.uber.org/zap"
	"io"
	"kgdt/configs"
	"xorm.io/xorm"
)

var (
	LogWriter   io.Writer
	Logger     	*zap.Logger
	DB         	*xorm.Engine
	Configs     *configs.Configs
	App *iris.Application
)


func init() {
	flag.Parse()
	//inits configs
	Configs = InitConfig()
	//inits log
	Logger,LogWriter = InitLogger(Configs.Log)
	//inits DB
	DB = InitDB(Configs.DB)
	//initApp
	App = InitApp()
}