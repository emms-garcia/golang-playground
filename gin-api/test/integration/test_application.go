package integration

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/gin-gonic/gin"
)

// TestApplication is a struct that holds the application and the gin engine for testing
type TestApplication struct {
	*application.Application
	Engine *gin.Engine
}

// NewTestApplication initializes a new TestApplication instance
func NewTestApplication() *TestApplication {
	app := application.Load()
	engine := router.Setup(app)
	return &TestApplication{
		Application: app,
		Engine:      engine,
	}
}

// Teardown performs cleanup operations after tests
func (a *TestApplication) Teardown() {
	tables := []string{"todos", "urls"}
	for _, table := range tables {
		result := a.DB.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY", table))
		if result.Error != nil {
			panic("failed to clear db")
		}
	}
}

// Request is a helper function to make HTTP requests to the test application
func (a *TestApplication) Request(method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	a.Engine.ServeHTTP(w, req)
	return w
}
