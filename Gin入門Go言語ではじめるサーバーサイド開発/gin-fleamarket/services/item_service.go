package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemSercice struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemSercice{repository: repository}
}

func (s *ItemSercice) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}
