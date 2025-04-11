package handler

import (
	"net/http"
	"strconv"

	"github.com/emms-garcia/golang-playground/gin-api/internal/service"
	"github.com/gin-gonic/gin"
)

// TodoHandler is an interface to define the methods for the todo handler
type TodoHandler interface {
	List(c *gin.Context)
	Get(c *gin.Context)
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type todoHandler struct {
	service service.TodoService
}

// NewTodoHandler is a function to create a new todo handler
func NewTodoHandler(service service.TodoService) TodoHandler {
	return &todoHandler{service: service}
}

// AddRequestBody is a struct to hold the request body (JSON) for the AddHandler
type AddRequestBody struct {
	Message string `json:"message"`
}

// Add is a handler function to add a todo
func (h *todoHandler) Add(c *gin.Context) {
	var requestBody AddRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	todo, err := h.service.CreateTodo(requestBody.Message)
	if err != nil {
		message := "unexpected error while creating todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// UpdateRequestBody is a struct to hold the request body (JSON) for the UpdateHandler
type UpdateRequestBody struct {
	Message string `json:"message"`
}

// Update is a handler function to update a todo by ID
func (h *todoHandler) Update(c *gin.Context) {
	var requestBody UpdateRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.UpdateTodo(id, requestBody.Message)
	if err != nil {
		message := "unexpected error while updating todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// List is a handler function to list all todos
func (h *todoHandler) List(c *gin.Context) {
	todos, err := h.service.GetTodos()
	if err != nil {
		message := "unexpected error while retrieving todos. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// Get is a handler function to get a todo by ID
func (h *todoHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.service.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// Delete is a handler function to delete a todo by ID
func (h *todoHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteTodo(id); err != nil {
		message := "unexpected error while deleting todo. please try again later"
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "todo deleted"})
}
