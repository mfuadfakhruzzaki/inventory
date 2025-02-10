package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"inventory-app/internal/models"
	"inventory-app/internal/services"
)

type TransactionHandler struct {
	Service services.TransactionService
}

func NewTransactionHandler(service services.TransactionService) *TransactionHandler {
	return &TransactionHandler{Service: service}
}

// GetTransactions mengembalikan transaksi untuk item tertentu.
// @Summary Ambil transaksi untuk item
// @Description Mengambil semua transaksi untuk item berdasarkan ID
// @Tags transactions
// @Param id path int true "ID Item"
// @Produce json
// @Success 200 {array} models.Transaction
// @Failure 400 {object} fiber.Map
// @Router /items/{id}/transactions [get]
func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID item tidak valid"})
	}
	transactions, err := h.Service.GetTransactionsByItemID(itemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(transactions)
}

// CreateTransaction membuat transaksi baru untuk item.
// @Summary Buat transaksi
// @Description Mencatat transaksi baru (stok masuk/keluar) untuk item tertentu
// @Tags transactions
// @Accept json
// @Produce json
// @Param id path int true "ID Item"
// @Param transaction body models.Transaction true "Data Transaksi"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} fiber.Map
// @Router /items/{id}/transactions [post]
func (h *TransactionHandler) CreateTransaction(c *fiber.Ctx) error {
	itemID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID item tidak valid"})
	}
	var transaction models.Transaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Input tidak valid"})
	}
	createdTransaction, err := h.Service.CreateTransaction(itemID, &transaction)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdTransaction)
}
