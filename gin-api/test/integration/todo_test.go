package integration

import (
	"net/http"
	"testing"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"github.com/stretchr/testify/assert"
)

// TestAddHandler tests the GET /todos endpoint
func TestAddHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	response := request(ctx, "POST", "/todos", "{\"message\": \"test\"}")
	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "{\"id\":1,\"message\":\"test\"}", response.Body.String())
}

// TestUpdateHandler tests the PUT /todos/:id endpoint
func TestUpdateHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	result := ctx.DB.Create(&model.Todo{Message: "test"})
	if result.Error != nil {
		t.Error("failed to create todo")
		return
	}

	response := request(ctx, "PUT", "/todos/1", "{\"message\": \"updated\"}")
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "{\"id\":1,\"message\":\"updated\"}", response.Body.String())
}

// TestDetailHandler tests the GET /todos/:id endpoint
func TestDetailHandler(t *testing.T) {
	ctx := setup()
	defer teardown(ctx)

	result := ctx.DB.Create(&model.Todo{Message: "test"})
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

	result1 := ctx.DB.Create(&model.Todo{Message: "test1"})
	result2 := ctx.DB.Create(&model.Todo{Message: "test2"})
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

	result1 := ctx.DB.Create(&model.Todo{Message: "test"})
	if result1.Error != nil {
		t.Error("failed to create todo")
		return
	}

	response := request(ctx, "DELETE", "/todos/1", "")
	assert.Equal(t, http.StatusNoContent, response.Code)
}
