package logger

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"go.uber.org/zap"
)

// Load initializes the logger based on the environment.
func Load(env string) (*zap.Logger, error) {
	if env == configuration.Production {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
