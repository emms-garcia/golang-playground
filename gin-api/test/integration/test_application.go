package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/emms-garcia/golang-playground/gin-api/internal/application"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/gin-gonic/gin"
)

type TestApplication struct {
	*application.Application
	Engine *gin.Engine
}

func NewTestApplication() *TestApplication {
	app := application.Load()
	engine := router.Setup(app)
	return &TestApplication{
		Application: app,
		Engine:      engine,
	}
}

func (a *TestApplication) Teardown() {
	result := a.DB.Exec("TRUNCATE TABLE todos RESTART IDENTITY")
	if result.Error != nil {
		panic("failed to clear db")
	}
}

func (a *TestApplication) Request(method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	a.Engine.ServeHTTP(w, req)
	return w
}
