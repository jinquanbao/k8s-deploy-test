package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"go.uber.org/zap"
	"kgdt/configs"
	"kgdt/inits"
	"kgdt/internal/intercept"
)

var (
	log *zap.SugaredLogger
)

func init() {
	log = inits.Logger.Sugar()
	Router(inits.App)
}

func Router(router iris.Party) {
	router.Get("/", Test)
}

// InitWeb 初始化web服务
func InitWeb(project configs.ProjectConfig) {
	inits.App.Use(intercept.AuthIntercept)
	inits.App.Use(func(ctx iris.Context) {
		ctx.Record()
		ctx.Next()
	})
	inits.App.Use(recover.New())
	inits.App.Use(intercept.LoggerIntercept)
	inits.App.Run(iris.Addr(":"+project.ServerPort), iris.WithCharset("utf-8"), iris.WithoutServerError(iris.ErrServerClosed))
}