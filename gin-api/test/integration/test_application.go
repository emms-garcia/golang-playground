package integration

import (
	"bytes"
	"context"
	"fmt"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/configuration"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/gin-gonic/gin"
)

// TestApplication is a struct that holds the application and the gin engine for testing
type TestApplication struct {
	*application.Application
	Engine *gin.Engine
}

// NewTestApplication initializes a new TestApplication instance
func NewTestApplication(t *testing.T) *TestApplication {
	t.Helper()

	app, err := application.Load(context.Background(), testConfigsPath(t), configuration.Test)
	if err != nil {
		t.Skipf("skipping integration test: %v", err)
	}

	engine := router.Setup(app)
	testApp := &TestApplication{
		Application: app,
		Engine:      engine,
	}
	t.Cleanup(func() {
		if err := testApp.Teardown(context.Background()); err != nil {
			t.Fatalf("teardown test app: %v", err)
		}

		if err := testApp.Close(); err != nil {
			t.Fatalf("close test app: %v", err)
		}
	})

	return testApp
}

// Teardown performs cleanup operations after tests
func (a *TestApplication) Teardown(ctx context.Context) error {
	tables := []string{"todos", "urls"}
	for _, table := range tables {
		result := a.DB.WithContext(ctx).Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if result.Error != nil {
			return fmt.Errorf("failed to clear %s: %w", table, result.Error)
		}
	}
	return nil
}

// Request is a helper function to make HTTP requests to the test application
func (a *TestApplication) Request(method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, "http://localhost"+path, bytes.NewBufferString(body))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	a.Engine.ServeHTTP(w, req)
	return w
}

func testConfigsPath(t *testing.T) string {
	t.Helper()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("failed to locate test file")
	}

	return filepath.Join(filepath.Dir(filename), "..", "..", "configs")
}
