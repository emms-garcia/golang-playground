package repository

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

// TodoRepository is an interface that defines the methods for the todo repository
type TodoRepository interface {
	GetTodos() ([]*model.Todo, error)
	GetTodoById(id int) (*model.Todo, error)
	Create(todo *model.Todo) error
	Update(todo *model.Todo) error
	Delete(todo *model.Todo) error
}

type todoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository is a function to create a new todo repository
func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{DB: db}
}

// GetTodoById is a function to get a todo by ID
func (r *todoRepository) GetTodoById(id int) (*model.Todo, error) {
	var todo model.Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// GetTodos is a function to get all todos
func (r *todoRepository) GetTodos() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Create is a method to create a new todo
func (r *todoRepository) Create(todo *model.Todo) error {
	return r.DB.Create(todo).Error
}

// Update is a method to update an existing todo
func (r *todoRepository) Update(todo *model.Todo) error {
	return r.DB.Save(todo).Error
}

// Delete is a method to delete a todo
func (r *todoRepository) Delete(todo *model.Todo) error {
	return r.DB.Delete(todo).Error
}
