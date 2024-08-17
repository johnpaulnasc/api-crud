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

func (r *ItemRepository) FindAll() ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *ItemRepository) Update(item *domain.Item) error {
	return r.db.Save(item).Error
}

func (r *ItemRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Item{}, id).Error
}