package repository

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

// TodoRepository is an interface that defines the methods for the todo repository
type TodoRepository interface {
	BaseRepository[model.Todo]
}

type todoRepository struct {
	BaseRepository[model.Todo]
	DB *gorm.DB
}

// NewTodoRepository is a function to create a new todo repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		BaseRepository: NewBaseRepository[model.Todo](db),
		DB:             db,
	}
}
