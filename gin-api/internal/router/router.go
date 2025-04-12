package router

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/handler"
	"github.com/emms-garcia/golang-playground/gin-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

// Setup is a function to set up the routes of the API
func Setup(handler *handler.Handler) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.SetTrustedProxies(nil)
	engine.Use(middleware.ZapLogger())

	// ping API (healthcheck)
	engine.GET("/ping", handler.PingHandler.Ping)

	// todos API
	engine.GET("/todos", handler.TodoHandler.List)
	engine.POST("/todos", handler.TodoHandler.Add)
	engine.GET("/todos/:id", handler.TodoHandler.Get)
	engine.PUT("/todos/:id", handler.TodoHandler.Update)
	engine.DELETE("/todos/:id", handler.TodoHandler.Delete)
	return engine
}
