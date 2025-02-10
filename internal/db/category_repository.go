package db

import (
	"github.com/jmoiron/sqlx"
	"inventory-app/internal/models"
)

type CategoryRepository interface {
	FetchAll() ([]models.Category, error)
	FetchByID(id int) (*models.Category, error)
	Insert(category *models.Category) (*models.Category, error)
	Update(category *models.Category) (*models.Category, error)
	Delete(id int) error
}

type categoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{DB: db}
}

func (r *categoryRepository) FetchAll() ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories`
	err := r.DB.Select(&categories, query)
	return categories, err
}

func (r *categoryRepository) FetchByID(id int) (*models.Category, error) {
	var category models.Category
	query := `SELECT * FROM categories WHERE id=$1`
	err := r.DB.Get(&category, query, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) Insert(category *models.Category) (*models.Category, error) {
	query := `INSERT INTO categories (name, description)
              VALUES (:name, :description)
              RETURNING id`
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	var id int
	if err := stmt.Get(&id, category); err != nil {
		return nil, err
	}
	category.ID = id
	return category, nil
}

func (r *categoryRepository) Update(category *models.Category) (*models.Category, error) {
	query := `UPDATE categories SET name=:name, description=:description WHERE id=:id`
	_, err := r.DB.NamedExec(query, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id int) error {
	query := `DELETE FROM categories WHERE id=$1`
	_, err := r.DB.Exec(query, id)
	return err
}
