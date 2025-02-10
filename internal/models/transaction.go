package models

import "time"

// Transaction merepresentasikan transaksi stok untuk item.
type Transaction struct {
	ID              int       `json:"id" db:"id"`
	ItemID          int       `json:"item_id" db:"item_id" validate:"required"`
	TransactionType string    `json:"transaction_type" db:"transaction_type" validate:"required,oneof=in out"`
	Quantity        int       `json:"quantity" db:"quantity" validate:"required,gt=0"`
	Timestamp       time.Time `json:"timestamp" db:"timestamp"`
	Notes           string    `json:"notes,omitempty" db:"notes"`
}
