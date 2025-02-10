package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"inventory-app/internal/models"
	"inventory-app/internal/services"
)

type ItemHandler struct {
	Service services.ItemService
}

func NewItemHandler(service services.ItemService) *ItemHandler {
	return &ItemHandler{Service: service}
}

// GetItems mengembalikan daftar item.
// @Summary Ambil daftar item
// @Description Mengambil semua item yang tersedia di inventaris
// @Tags items
// @Produce json
// @Success 200 {array} models.Item
// @Router /items [get]
func (h *ItemHandler) GetItems(c *fiber.Ctx) error {
	items, err := h.Service.GetAllItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(items)
}

// GetItem mengembalikan detail item berdasarkan ID.
// @Summary Ambil detail item
// @Description Mengambil detail item berdasarkan ID yang diberikan
// @Tags items
// @Param id path int true "ID Item"
// @Produce json
// @Success 200 {object} models.Item
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /items/{id} [get]
func (h *ItemHandler) GetItem(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID tidak valid"})
	}
	item, err := h.Service.GetItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "Item tidak ditemukan"})
	}
	return c.JSON(item)
}

// CreateItem membuat item baru.
// @Summary Buat item baru
// @Description Menambahkan item baru ke dalam inventaris
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.Item true "Data Item"
// @Success 201 {object} models.Item
// @Failure 400 {object} fiber.Map
// @Router /items [post]
func (h *ItemHandler) CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Input tidak valid"})
	}
	createdItem, err := h.Service.CreateItem(&item)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdItem)
}
