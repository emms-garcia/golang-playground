package logger

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/config"
	"go.uber.org/zap"
)

var Log *zap.Logger

func Init(env string) {
	var err error
	if env == config.Production {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
}
