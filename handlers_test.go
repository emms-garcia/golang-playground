package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// TestContext is a struct to hold the context used in tests
type TestContext struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

// setup is a helper function to run before each test (i.e. to set up the database)
func setup() *TestContext {
	db := ConfigureDB(&Configuration{
		DBHost:     "testdb",
		DBUser:     "postgres",
		DBPassword: "123456",
		DBName:     "db",
	})
	engine := ConfigureRoutes(db)
	return &TestContext{DB: db, Engine: engine}
}

// teardown is a helper function to run after each test (i.e. to clean up the database)
func teardown(ctx *TestContext) {
	result := ctx.DB.Exec("TRUNCATE TABLE todos RESTART IDENTITY")
	if result.Error != nil {
		panic("failed to clear db")
	}
}

// request is a helper function to make a request to the server
func request(ctx *TestContext, method, path, body string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, bytes.NewBufferString(body))
	ctx.Engine.ServeHTTP(w, req)
	return w
}

// TestPingHandler tests the GET /ping endpoint
func TestPingHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	response := request(ctx, "GET", "/ping", "")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", response.Body.String())
}

// TestAddHandler tests the GET /todos endpoint
func TestAddHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	response := request(ctx, "POST", "/todos", "{\"message\": \"test\"}")
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "{\"id\":1,\"message\":\"test\"}", response.Body.String())
}

// TestDetailHandler tests the GET /todos/:id endpoint
func TestDetailHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	result := ctx.DB.Create(&Todo{Message: "test"})
	if result.Error != nil {
		t.Error("failed to create todo")
		return
	}

	response := request(ctx, "GET", "/todos/1", "")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"id\":1,\"message\":\"test\"}", response.Body.String())
}

// TestListHandler tests the GET /todos endpoint
func TestListHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	result1 := ctx.DB.Create(&Todo{Message: "test1"})
	result2 := ctx.DB.Create(&Todo{Message: "test2"})
	if result1.Error != nil || result2.Error != nil {
		t.Error("failed to create todos")
		return
	}

	response := request(ctx, "GET", "/todos", "")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "[{\"id\":1,\"message\":\"test1\"},{\"id\":2,\"message\":\"test2\"}]", response.Body.String())
}

// TestDeleteHandler tests the DELETE /todos/:id endpoint
func TestDeleteHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	result1 := ctx.DB.Create(&Todo{Message: "test"})
	if result1.Error != nil {
		t.Error("failed to create todo")
		return
	}

	response := request(ctx, "DELETE", "/todos/1", "")
	assert.Equal(t, http.StatusNoContent, response.Code)
}
