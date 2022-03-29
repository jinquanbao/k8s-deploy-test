package controller

import "github.com/kataras/iris/v12"

func Test(ctx iris.Context) {
	log.Info("requert uri =",ctx.Request().RequestURI)
	ctx.Text("request success ...")
}