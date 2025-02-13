package main

// Todo is a struct to hold the model for the "todos" table
type Todo struct {
	ID      int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Message string `json:"message"`
}
