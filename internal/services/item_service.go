package services

import (
	"inventory-app/internal/db"
	"inventory-app/internal/models"

	"github.com/go-playground/validator/v10"
)

// ItemService interface untuk logika bisnis entitas Item.
type ItemService interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id int) (*models.Item, error)
	CreateItem(item *models.Item) (*models.Item, error)
	// Tambahkan fungsi UpdateItem dan DeleteItem bila diperlukan.
}

type itemService struct {
	repo      db.ItemRepository
	validator *validator.Validate
}

// NewItemService mengembalikan instance ItemService.
func NewItemService(repo db.ItemRepository) ItemService {
	return &itemService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *itemService) GetAllItems() ([]models.Item, error) {
	return s.repo.FetchAll()
}

func (s *itemService) GetItemByID(id int) (*models.Item, error) {
	return s.repo.FetchByID(id)
}

func (s *itemService) CreateItem(item *models.Item) (*models.Item, error) {
	// Validasi input
	if err := s.validator.Struct(item); err != nil {
		return nil, err
	}
	return s.repo.Insert(item)
}
