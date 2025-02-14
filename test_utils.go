package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TestContext is a struct to hold the context used in tests
type TestContext struct {
	DB     *gorm.DB
	Engine *gin.Engine
}

var TestConfiguration Configuration = Configuration{
	DBHost:     "testdb",
	DBUser:     "postgres",
	DBPassword: "123456",
	DBName:     "db",
}

// setup is a helper function to run before each test (i.e. to set up the database)
func setup() *TestContext {
	db := ConfigureDB(&TestConfiguration)
	app := NewApp(db)
	engine := ConfigureRoutes(app)
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
