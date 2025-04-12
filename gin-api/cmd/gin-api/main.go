package main

import (
	"fmt"
	"log"
	"os"

	"github.com/emms-garcia/golang-playground/gin-api/internal/config"
	"github.com/emms-garcia/golang-playground/gin-api/internal/handler"
	"github.com/emms-garcia/golang-playground/gin-api/internal/logger"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	// Load configuration
	if err := config.Load(environment); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Initialize logger
	logger.Init(environment)
	// Initialize database connection
	db := config.ConfigureDB(&config.AppConfiguration)
	// Initialize handlers
	handler := &handler.Handler{
		PingHandler: handler.NewPingHandler(),
		TodoHandler: handler.NewTodoHandler(service.NewTodoService(repository.NewTodoRepository(db))),
	}
	// Initialize router
	engine := router.Setup(handler)
	engine.Run(fmt.Sprintf(":%d", config.AppConfiguration.Server.Port))
}
