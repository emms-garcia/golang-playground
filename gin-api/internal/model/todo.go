package model

// Todo is a struct that represents a row in the "todos" table
type Todo struct {
	ID      int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Message string `json:"message"`
}
