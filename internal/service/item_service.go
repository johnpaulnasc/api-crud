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

// Outros métodos: FindAll, Update, Delete...