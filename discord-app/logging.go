package main

import "go.uber.org/zap"

func InitLogger(environment string) *zap.Logger {
	var logger *zap.Logger
	if IsProduction(environment) {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)
	return logger
}
