package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/emms-garcia/golang-playground/gin-api/internal/config"
	"github.com/emms-garcia/golang-playground/gin-api/internal/handler"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
	"github.com/emms-garcia/golang-playground/gin-api/internal/router"
	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IntegrationTestContext struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

func setup() *IntegrationTestContext {
	db := config.ConfigureDB(&config.TestConfiguration)
	engine := router.Setup(&handler.Handler{
		PingHandler: handler.NewPingHandler(),
		TodoHandler: handler.NewTodoHandler(service.NewTodoService(repository.NewTodoRepository(db))),
	})
	return &IntegrationTestContext{
		DB:     db,
		Engine: engine,
	}
}

func teardown(ctx *IntegrationTestContext) {
	result := ctx.DB.Exec("TRUNCATE TABLE todos RESTART IDENTITY")
	if result.Error != nil {
		panic("failed to clear db")
	}
}

func request(ctx *IntegrationTestContext, method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	ctx.Engine.ServeHTTP(w, req)
	return w
}
