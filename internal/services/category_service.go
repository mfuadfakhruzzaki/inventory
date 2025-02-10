package services

import (
	"github.com/go-playground/validator/v10"
	"inventory-app/internal/db"
	"inventory-app/internal/models"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryByID(id int) (*models.Category, error)
	CreateCategory(category *models.Category) (*models.Category, error)
	UpdateCategory(category *models.Category) (*models.Category, error)
	DeleteCategory(id int) error
}

type categoryService struct {
	repo      db.CategoryRepository
	validator *validator.Validate
}

func NewCategoryService(repo db.CategoryRepository) CategoryService {
	return &categoryService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.FetchAll()
}

func (s *categoryService) GetCategoryByID(id int) (*models.Category, error) {
	return s.repo.FetchByID(id)
}

func (s *categoryService) CreateCategory(category *models.Category) (*models.Category, error) {
	if err := s.validator.Struct(category); err != nil {
		return nil, err
	}
	return s.repo.Insert(category)
}

func (s *categoryService) UpdateCategory(category *models.Category) (*models.Category, error) {
	if err := s.validator.Struct(category); err != nil {
		return nil, err
	}
	return s.repo.Update(category)
}

func (s *categoryService) DeleteCategory(id int) error {
	return s.repo.Delete(id)
}
