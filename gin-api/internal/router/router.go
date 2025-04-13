package router

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/handler"
	"github.com/emms-garcia/golang-playground/gin-api/internal/middleware"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
)

// Setup is a function to set up the routes of the API
func Setup(app *application.Application) *gin.Engine {
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(gin.Recovery())
	engine.Use(middleware.LogRequestMiddleware(app.Logger))

	// ping API (healthcheck)
	pingHandler := handler.NewPingHandler()
	engine.GET("/ping", pingHandler.Ping)

	// todos API
	todoHandler := handler.NewTodoHandler(service.NewTodoService(repository.NewTodoRepository(app.DB)))
	engine.GET("/todos", todoHandler.List)
	engine.POST("/todos", todoHandler.Add)
	engine.GET("/todos/:id", todoHandler.Get)
	engine.PUT("/todos/:id", todoHandler.Update)
	engine.DELETE("/todos/:id", todoHandler.Delete)
	return engine
}
