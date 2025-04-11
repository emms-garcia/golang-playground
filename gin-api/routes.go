package main

import (
	"github.com/gin-gonic/gin"
)

// ConfigureRoutes is a function to configure the routes of the API
func ConfigureRoutes(app *App) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.SetTrustedProxies(nil)
	engine.GET("/ping", app.PingHandler)
	engine.GET("/todos", app.ListHandler)
	engine.POST("/todos", app.AddHandler)
	engine.GET("/todos/:id", app.DetailHandler)
	engine.PUT("/todos/:id", app.UpdateHandler)
	engine.DELETE("/todos/:id", app.DeleteHandler)
	return engine
}
