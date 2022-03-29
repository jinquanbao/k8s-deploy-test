package inits

import (
	"github.com/kataras/iris/v12"
)


func InitApp() *iris.Application{
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Logger().SetOutput(LogWriter)

	return app
}


