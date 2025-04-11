package service

import (
	"github.com/emms-garcia/golang-playground/gin-api/internal/model"
	"github.com/emms-garcia/golang-playground/gin-api/internal/repository"
)

// TodoService is an interface that defines the methods for the todo service
type TodoService interface {
	GetTodos() ([]*model.Todo, error)
	GetTodo(id int) (*model.Todo, error)
	CreateTodo(message string) (*model.Todo, error)
	UpdateTodo(id int, message string) (*model.Todo, error)
	DeleteTodo(id int) error
}

type todoService struct {
	repo repository.TodoRepository
}

// NewTodoService is a function to create a new todo service
func NewTodoService(repo repository.TodoRepository) TodoService {
	return &todoService{repo: repo}
}

// CreateTodo is a method to create a new todo
// It takes a message as input and returns the created Todo or an error
func (s *todoService) CreateTodo(message string) (*model.Todo, error) {
	todo := &model.Todo{Message: message}
	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// UpdateTodo is a method to update an existing todo
// It takes an ID and a message as input and returns the updated Todo or an error
func (s *todoService) UpdateTodo(id int, message string) (*model.Todo, error) {
	todo, err := s.repo.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	todo.Message = message
	if err = s.repo.Update(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

// GetTodos is a method to get all todos
// It returns a slice of Todos or an error
func (s *todoService) GetTodos() ([]*model.Todo, error) {
	todos, err := s.repo.GetTodos()
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// GetTodo is a method to get a todo by ID
// It takes an ID as input and returns the Todo or an error
func (s *todoService) GetTodo(id int) (*model.Todo, error) {
	todo, err := s.repo.GetTodoById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// DeleteTodo is a method to delete a todo by ID
// It takes an ID as input and returns nil if the deletion was successful or an error
func (s *todoService) DeleteTodo(id int) error {
	todo, err := s.GetTodo(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(todo)
}
