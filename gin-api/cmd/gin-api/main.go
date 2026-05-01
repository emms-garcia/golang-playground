package main

import (
	"context"
	"fmt"
	"log"

	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
)

func main() {
	// Get the configuration path
	configsPath, err := configuration.GetConfigsPath()
	if err != nil {
		log.Fatal(err)
	}

	// Get the environment
	environment, err := configuration.GetEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// Load the application
	app, err := application.Load(context.Background(), configsPath, environment)
	if err != nil {
		log.Fatal(err)
	}
	defer app.Close()

	// Setup the router
	engine := router.Setup(app)
	if err := engine.Run(fmt.Sprintf(":%d", app.Config.Server.Port)); err != nil {
		err = fmt.Errorf("run server: %w", err)
		log.Fatal(err)
	}
}
