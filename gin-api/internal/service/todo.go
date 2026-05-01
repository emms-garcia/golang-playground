package service

import (
	"context"

	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

// NewTodoService is a function to create a new todo service
func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

// CreateTodo is a method to create a new todo
// It takes a message as input and returns the created Todo or an error
func (s *TodoService) CreateTodo(ctx context.Context, message string) (*model.Todo, error) {
	todo := &model.Todo{Message: message}
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// UpdateTodo is a method to update an existing todo
// It takes an ID and a message as input and returns the updated Todo or an error
func (s *TodoService) UpdateTodo(ctx context.Context, id int, message string) (*model.Todo, error) {
	todo, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.Message = message
	if err = s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// GetTodos is a method to get all todos
// It returns a slice of Todos or an error
func (s *TodoService) GetTodos(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.All(ctx)
}

// GetTodo is a method to get a todo by ID
// It takes an ID as input and returns the Todo or an error
func (s *TodoService) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	return s.repo.Get(ctx, id)
}

// DeleteTodo is a method to delete a todo by ID
// It takes an ID as input and returns nil if the deletion was successful or an error
func (s *TodoService) DeleteTodo(ctx context.Context, id int) error {
	todo, err := s.GetTodo(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, todo)
}
