package intercept

import "github.com/kataras/iris/v12"

func AuthIntercept(ctx iris.Context) {
	log.Info("Auth Intercept ........")
}
