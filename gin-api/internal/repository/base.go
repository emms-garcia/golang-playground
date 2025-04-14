package repository

import "gorm.io/gorm"

// BaseRepository is a generic interface for basic CRUD operations
type BaseRepository[T any] interface {
	All() ([]*T, error)
	Get(id int) (*T, error)
	Create(item *T) error
	Update(item *T) error
	Delete(item *T) error
}

type baseRepository[T any] struct {
	Db *gorm.DB
}

// NewBaseRepository is a function to create a new base repository
func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{Db: db}
}

// All is a method to get all items
func (r *baseRepository[T]) All() ([]*T, error) {
	var items []*T
	if err := r.Db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

// Get is a method to get an item by ID
func (r *baseRepository[T]) Get(id int) (*T, error) {
	var item T
	if err := r.Db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// Create is a method to create a new item
func (r *baseRepository[T]) Create(item *T) error {
	return r.Db.Create(item).Error
}

// Update is a method to update an item
func (r *baseRepository[T]) Update(item *T) error {
	return r.Db.Save(item).Error
}

// Delete is a method to delete an item
func (r *baseRepository[T]) Delete(item *T) error {
	return r.Db.Delete(item).Error
}
