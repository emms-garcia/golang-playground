package repository

import (
	"context"

	"gorm.io/gorm"
)

type baseRepository[T any] struct {
	db *gorm.DB
}

func newBaseRepository[T any](db *gorm.DB) *baseRepository[T] {
	return &baseRepository[T]{db: db}
}

// All is a method to get all items
func (r *baseRepository[T]) All(ctx context.Context) ([]*T, error) {
	var items []*T
	if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

// Get is a method to get an item by ID
func (r *baseRepository[T]) Get(ctx context.Context, id int) (*T, error) {
	var item T
	if err := r.db.WithContext(ctx).First(&item, id).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

// Create is a method to create a new item
func (r *baseRepository[T]) Create(ctx context.Context, item *T) error {
	return r.db.WithContext(ctx).Create(item).Error
}

// Update is a method to update an item
func (r *baseRepository[T]) Update(ctx context.Context, item *T) error {
	return r.db.WithContext(ctx).Save(item).Error
}

// Delete is a method to delete an item
func (r *baseRepository[T]) Delete(ctx context.Context, item *T) error {
	return r.db.WithContext(ctx).Delete(item).Error
}
