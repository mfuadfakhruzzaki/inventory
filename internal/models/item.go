package models

// Item merepresentasikan entitas barang.
type Item struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" validate:"required"`
	Description string `json:"description,omitempty" db:"description"`
	Quantity    int    `json:"quantity" db:"quantity" validate:"gte=0"`
	Location    string `json:"location" db:"location" validate:"required"`
	CategoryID  int    `json:"category_id" db:"category_id" validate:"required"`
}
