package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, v := range r.items {
		if v.ID == updateItem.ID {
			r.items[i] = updateItem
			return &r.items[i], nil
		}
	}
	return nil, errors.New("unexpected error")
}

func (r *ItemMemoryRepository) Delete(itemId uint) error {
	for i, v := range r.items {
		if v.ID == itemId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

// Create implements IItemRepository.
func (r *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := r.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (i *ItemRepository) Delete(itemId uint) error {
	panic("unimplemented")
}

// FindAll implements IItemRepository.
func (r *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := r.db.Find(&items)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &items, nil
}

// FindById implements IItemRepository.
func (r *ItemRepository) FindById(itemId uint) (*models.Item, error) {
	var item models.Item
	result := r.db.First(&item, itemId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

// Update implements IItemRepository.
func (r *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	// Save: upsert処理
	result := r.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}
