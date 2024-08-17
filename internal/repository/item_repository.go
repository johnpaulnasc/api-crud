package repository

import (
    "gorm.io/gorm"
    "api-crud/internal/domain"
)

type ItemRepository struct {
    db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
    return &ItemRepository{db: db}
}

func (r *ItemRepository) Create(item *domain.Item) error {
    return r.db.Create(item).Error
}

// Outros métodos: FindAll, Update, Delete...