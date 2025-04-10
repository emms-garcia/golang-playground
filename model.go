package main

import "gorm.io/gorm"

// Todo is a struct to hold the model for the "todos" table
type Todo struct {
	ID      int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Message string `json:"message"`
}

// GetTodoById is a function to get a todo by ID
func GetTodoById(db *gorm.DB, id int) (*Todo, error) {
	var todo Todo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// GetTodos is a function to get all todos
func GetTodos(db *gorm.DB) ([]*Todo, error) {
	var todos []*Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Create is a method to create a new todo
func (t *Todo) Create(db *gorm.DB) error {
	todo := db.Create(t)
	return todo.Error
}

// Update is a method to update an existing todo
func (t *Todo) Update(db *gorm.DB) error {
	todo := db.Save(t)
	return todo.Error
}

// Delete is a method to delete a todo
func (t *Todo) Delete(db *gorm.DB) error {
	todo := db.Delete(t)
	return todo.Error
}
