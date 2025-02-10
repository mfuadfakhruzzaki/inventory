package services

import (
	"time"

	"github.com/go-playground/validator/v10"
	"inventory-app/internal/db"
	"inventory-app/internal/models"
)

type TransactionService interface {
	GetTransactionsByItemID(itemID int) ([]models.Transaction, error)
	CreateTransaction(itemID int, transaction *models.Transaction) (*models.Transaction, error)
}

type transactionService struct {
	repo      db.TransactionRepository
	validator *validator.Validate
}

func NewTransactionService(repo db.TransactionRepository) TransactionService {
	return &transactionService{
		repo:      repo,
		validator: validator.New(),
	}
}

func (s *transactionService) GetTransactionsByItemID(itemID int) ([]models.Transaction, error) {
	return s.repo.FetchByItemID(itemID)
}

func (s *transactionService) CreateTransaction(itemID int, transaction *models.Transaction) (*models.Transaction, error) {
	transaction.ItemID = itemID
	if transaction.Timestamp.IsZero() {
		transaction.Timestamp = time.Now()
	}
	if err := s.validator.Struct(transaction); err != nil {
		return nil, err
	}
	return s.repo.Insert(transaction)
}
