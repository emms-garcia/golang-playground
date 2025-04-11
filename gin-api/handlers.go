package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// App is a struct to hold the application context
type App struct {
	DB *gorm.DB
}

// NewApp is a function to create a new App instance
func NewApp(db *gorm.DB) *App {
	return &App{DB: db}
}

// PingHandler is a handler function to respond to the ping endpoint
func (app *App) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// AddRequestBody is a struct to hold the request body (JSON) for the AddHandler
type AddRequestBody struct {
	Message string `json:"message"`
}

// AddHandler is a handler function to add a todo
func (app *App) AddHandler(c *gin.Context) {
	var requestBody AddRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	todo := &Todo{Message: requestBody.Message}
	if err := todo.Create(app.DB); err != nil {
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

// UpdateHandler is a handler function to update a todo by ID
func (app *App) UpdateHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := GetTodoById(app.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	var requestBody UpdateRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	todo.Message = requestBody.Message
	if err := todo.Update(app.DB); err != nil {
		message := "unexpected error while updating todo. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// ListHandler is a handler function to list all todos
func (app *App) ListHandler(c *gin.Context) {
	todos, err := GetTodos(app.DB)
	if err != nil {
		message := "unexpected error while retrieving todos. please try again later"
		c.JSON(http.StatusInternalServerError, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// DetailHandler is a handler function to get a todo by ID
func (app *App) DetailHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := GetTodoById(app.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteHandler is a handler function to delete a todo by ID
func (app *App) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := GetTodoById(app.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
		return
	}
	if err = todo.Delete(app.DB); err != nil {
		message := "unexpected error while deleting todo. please try again later"
		c.JSON(http.StatusNotFound, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "todo deleted"})
}
