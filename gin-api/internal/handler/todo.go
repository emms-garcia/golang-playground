package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TodoHandler struct holds the todo service instance.
type TodoHandler struct {
	service *service.TodoService
}

// NewTodoHandler is a function to create a new todo handler
func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

// AddRequestBody is a struct to hold the request body (JSON) for the AddHandler
type AddRequestBody struct {
	Message string `json:"message" binding:"required"`
}

// Add is a handler function to add a todo
func (h *TodoHandler) Add(c *gin.Context) {
	var requestBody AddRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	todo, err := h.service.CreateTodo(c.Request.Context(), requestBody.Message)
	if err != nil {
		message := "unexpected error while creating todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// UpdateRequestBody is a struct to hold the request body (JSON) for the UpdateHandler
type UpdateRequestBody struct {
	Message string `json:"message" binding:"required"`
}

// Update is a handler function to update a todo by ID
func (h *TodoHandler) Update(c *gin.Context) {
	var requestBody UpdateRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	id, err := parseTodoID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	todo, err := h.service.UpdateTodo(c.Request.Context(), id, requestBody.Message)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		message := "unexpected error while updating todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// List is a handler function to list all todos
func (h *TodoHandler) List(c *gin.Context) {
	todos, err := h.service.GetTodos(c.Request.Context())
	if err != nil {
		message := "unexpected error while retrieving todos. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// Get is a handler function to get a todo by ID
func (h *TodoHandler) Get(c *gin.Context) {
	id, err := parseTodoID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	todo, err := h.service.GetTodo(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected error while retrieving todo. please try again later"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// Delete is a handler function to delete a todo by ID
func (h *TodoHandler) Delete(c *gin.Context) {
	id, err := parseTodoID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid todo id"})
		return
	}

	if err := h.service.DeleteTodo(c.Request.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}

		message := "unexpected error while deleting todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}

	c.Status(http.StatusNoContent)
}

func parseTodoID(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		return 0, errors.New("invalid todo id")
	}

	return id, nil
}
