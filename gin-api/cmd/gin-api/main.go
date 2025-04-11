package main

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/config"
	"github.com/emms-garcia/golang-playground/gin-api/internal/handler"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
)

func main() {
	db := config.ConfigureDB(&config.Configuration{
		DBHost:     "db",
		DBUser:     "postgres",
		DBPassword: "123456",
		DBName:     "postgres",
	})
	handler := &handler.Handler{
		PingHandler: handler.NewPingHandler(),
		TodoHandler: handler.NewTodoHandler(service.NewTodoService(repository.NewTodoRepository(db))),
	}
	engine := router.Setup(handler)
	engine.Run(":8080")
}
