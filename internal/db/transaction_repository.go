package db

import (
	"github.com/jmoiron/sqlx"
	"inventory-app/internal/models"
)

type TransactionRepository interface {
	FetchByItemID(itemID int) ([]models.Transaction, error)
	Insert(transaction *models.Transaction) (*models.Transaction, error)
}

type transactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{DB: db}
}

func (r *transactionRepository) FetchByItemID(itemID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := `SELECT * FROM transactions WHERE item_id=$1 ORDER BY timestamp DESC`
	err := r.DB.Select(&transactions, query, itemID)
	return transactions, err
}

func (r *transactionRepository) Insert(transaction *models.Transaction) (*models.Transaction, error) {
	query := `INSERT INTO transactions (item_id, transaction_type, quantity, timestamp, notes)
              VALUES (:item_id, :transaction_type, :quantity, :timestamp, :notes)
              RETURNING id`
	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	var id int
	if err := stmt.Get(&id, transaction); err != nil {
		return nil, err
	}
	transaction.ID = id
	return transaction, nil
}
