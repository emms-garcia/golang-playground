package main

import (
	"fmt"

	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
)

func main() {
	// Load the application
	app := application.Load()
	// Setup the router
	engine := router.Setup(app)
	engine.Run(fmt.Sprintf(":%d", app.Config.Server.Port))
}
