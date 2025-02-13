package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// PingHandler is a handler function to simply check if the API is running
func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// AddRequestBody is a struct to hold the request body (JSON) for the AddHandler
type AddRequestBody struct {
	Message string `json:"message"`
}

// AddHandler is a handler function to add a new todo
func AddHandler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var requestBody AddRequestBody
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		todo := &Todo{Message: requestBody.Message}
		result := db.Create(&todo)
		if result.Error != nil {
			message := "unexpected error while creating todo. please try again later"
			c.JSON(http.StatusInternalServerError, gin.H{"error": message})
			return
		}
		c.JSON(http.StatusCreated, todo)
	}
}

// ListHandler is a handler function to list all todos
func ListHandler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		todos := []Todo{}
		result := db.Find(&todos)
		if result.Error != nil {
			message := "unexpected error while retrieving todos. please try again later"
			c.JSON(http.StatusInternalServerError, gin.H{"error": message})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}

// DetailHandler is a handler function to get a todo by ID
func DetailHandler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		todo := Todo{ID: id}
		result := db.First(&todo)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteHandler is a handler function to delete a todo by ID
func DeleteHandler(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		todo := Todo{ID: id}
		result := db.First(&todo)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
			return
		}
		result = db.Delete(&todo)
		if result.Error != nil {
			message := "unexpected error while deleting todo. please try again later"
			c.JSON(http.StatusNotFound, gin.H{"error": message})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "todo deleted"})
	}
}
