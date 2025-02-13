package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ConfigureRoutes is a function to configure the routes of the API
func ConfigureRoutes(db *gorm.DB) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.GET("/ping", PingHandler)
	engine.GET("/todos", ListHandler(db))
	engine.POST("/todos", AddHandler(db))
	engine.GET("/todos/:id", DetailHandler(db))
	engine.DELETE("/todos/:id", DeleteHandler(db))
	return engine
}
