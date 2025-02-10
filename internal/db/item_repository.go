package db

import (
	"inventory-app/internal/models"

	"github.com/jmoiron/sqlx"
)

// ItemRepository interface untuk operasi database pada entitas Item.
type ItemRepository interface {
	FetchAll() ([]models.Item, error)
	FetchByID(id int) (*models.Item, error)
	Insert(item *models.Item) (*models.Item, error)
	// Tambahkan fungsi Update dan Delete bila diperlukan.
}

type itemRepository struct {
	DB *sqlx.DB
}

// NewItemRepository mengembalikan instance ItemRepository.
func NewItemRepository(db *sqlx.DB) ItemRepository {
	return &itemRepository{DB: db}
}

func (r *itemRepository) FetchAll() ([]models.Item, error) {
	var items []models.Item
	query := `SELECT * FROM items`
	err := r.DB.Select(&items, query)
	return items, err
}

func (r *itemRepository) FetchByID(id int) (*models.Item, error) {
	var item models.Item
	query := `SELECT * FROM items WHERE id=$1`
	err := r.DB.Get(&item, query, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Insert(item *models.Item) (*models.Item, error) {
	query := `INSERT INTO items (name, description, quantity, location, category_id)
              VALUES (:name, :description, :quantity, :location, :category_id) RETURNING id`
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	var id int
	if err := stmt.Get(&id, item); err != nil {
		return nil, err
	}
	item.ID = id
	return item, nil
}
