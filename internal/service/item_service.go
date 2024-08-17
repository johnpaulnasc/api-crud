package service

import (
	"api-crud/internal/domain"
	"api-crud/internal/repository"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item *domain.Item) error {
	return s.repo.Create(item)
}

func (s *ItemService) GetAllItems() ([]domain.Item, error) {
	return s.repo.FindAll()
}

func (s *ItemService) UpdateItem(item *domain.Item) error {
	return s.repo.Update(item)
}

func (s *ItemService) DeleteItem(id uint) error {
	return s.repo.Delete(id)
}