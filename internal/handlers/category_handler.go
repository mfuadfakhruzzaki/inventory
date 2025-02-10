package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"inventory-app/internal/models"
	"inventory-app/internal/services"
)

type CategoryHandler struct {
	Service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{Service: service}
}

// GetCategories mengembalikan daftar kategori.
// @Summary Ambil semua kategori
// @Description Mengambil daftar semua kategori yang ada
// @Tags categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(c *fiber.Ctx) error {
	categories, err := h.Service.GetAllCategories()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(categories)
}

// GetCategory mengembalikan detail kategori berdasarkan ID.
// @Summary Ambil detail kategori
// @Description Mengambil detail kategori berdasarkan ID
// @Tags categories
// @Param id path int true "ID Kategori"
// @Produce json
// @Success 200 {object} models.Category
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID tidak valid"})
	}
	category, err := h.Service.GetCategoryByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": "Kategori tidak ditemukan"})
	}
	return c.JSON(category)
}

// CreateCategory membuat kategori baru.
// @Summary Buat kategori baru
// @Description Menambahkan kategori baru
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Data Kategori"
// @Success 201 {object} models.Category
// @Failure 400 {object} fiber.Map
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Input tidak valid"})
	}
	createdCategory, err := h.Service.CreateCategory(&category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(createdCategory)
}

// UpdateCategory memperbarui kategori.
// @Summary Update kategori
// @Description Memperbarui kategori berdasarkan ID dan data yang diberikan
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "ID Kategori"
// @Param category body models.Category true "Data Kategori"
// @Success 200 {object} models.Category
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID tidak valid"})
	}
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Input tidak valid"})
	}
	category.ID = id
	updatedCategory, err := h.Service.UpdateCategory(&category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedCategory)
}

// DeleteCategory menghapus kategori berdasarkan ID.
// @Summary Hapus kategori
// @Description Menghapus kategori berdasarkan ID
// @Tags categories
// @Param id path int true "ID Kategori"
// @Success 204
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "ID tidak valid"})
	}
	if err := h.Service.DeleteCategory(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
