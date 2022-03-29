package intercept

import (
	"go.uber.org/zap"
	"kgdt/inits"
)

var (
	log *zap.SugaredLogger
)

func init() {
	log = inits.Logger.Sugar()
}