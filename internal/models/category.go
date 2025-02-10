package models

// Category merepresentasikan kategori barang.
type Category struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description,omitempty" db:"description"`
}
