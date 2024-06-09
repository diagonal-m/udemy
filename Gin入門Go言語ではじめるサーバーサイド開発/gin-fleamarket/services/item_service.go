package services

import (
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
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

func (s *ItemSercice) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}
