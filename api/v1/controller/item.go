package controller

import (
    "encoding/json"
    "net/http"
    "api-crud/api/v1/dto"
    "api-crud/internal/domain"
    "api-crud/internal/service"
)

type ItemController struct {
    service *service.ItemService
}

func NewItemController(service *service.ItemService) *ItemController {
    return &ItemController{service: service}
}

func (c *ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {
    var itemDTO dto.ItemDTO
    json.NewDecoder(r.Body).Decode(&itemDTO)

    item := &domain.Item{
        Name:  itemDTO.Name,
        Price: itemDTO.Price,
    }

    err := c.service.CreateItem(item)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

// Outros métodos: FindAll, Update, Delete...
