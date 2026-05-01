package repository

import (
	"context"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"gorm.io/gorm"
)

// TodoRepository struct holds the base repository for the Todo model.
type TodoRepository struct {
	*baseRepository[model.Todo]
}

// NewTodoRepository is a function to create a new todo repository
func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{baseRepository: newBaseRepository[model.Todo](db)}
}

// All is a method to get all todos
func (r *TodoRepository) All(ctx context.Context) ([]*model.Todo, error) {
	return r.baseRepository.All(ctx)
}

// Get is a method to get a todo by ID
func (r *TodoRepository) Get(ctx context.Context, id int) (*model.Todo, error) {
	return r.baseRepository.Get(ctx, id)
}

// Create is a method to create a new todo
func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	return r.baseRepository.Create(ctx, todo)
}

// Update is a method to update a todo
func (r *TodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	return r.baseRepository.Update(ctx, todo)
}

// Delete is a method to delete a todo
func (r *TodoRepository) Delete(ctx context.Context, todo *model.Todo) error {
	return r.baseRepository.Delete(ctx, todo)
}
