package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"inventory-app/internal/db"
	"inventory-app/internal/handlers"
	"inventory-app/internal/services"
)

func main() {
	// Ambil DSN dari environment variable, gunakan default jika tidak ditemukan.
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:020803@db:5432/inventory?sslmode=disable"
	}

	// Tambahkan logika retry untuk menunggu sampai database siap
	var dbConn *sqlx.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		dbConn, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		log.Printf("Database belum siap: %v. Mencoba lagi dalam 2 detik... (%d/%d)", err, i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Gagal terkoneksi ke database setelah %d percobaan: %v", maxRetries, err)
	}

	// Inisialisasi repository, service, dan handler untuk Items
	itemRepo := db.NewItemRepository(dbConn)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	// Inisialisasi repository, service, dan handler untuk Categories
	categoryRepo := db.NewCategoryRepository(dbConn)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Inisialisasi repository, service, dan handler untuk Transactions
	transactionRepo := db.NewTransactionRepository(dbConn)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Inisialisasi Fiber app
	app := fiber.New()
	api := app.Group("/api")

	// Endpoint untuk Items
	items := api.Group("/items")
	items.Get("/", itemHandler.GetItems)
	items.Get("/:id", itemHandler.GetItem)
	items.Post("/", itemHandler.CreateItem)
	// Endpoint PUT dan DELETE untuk items bisa ditambahkan dengan pola yang sama

	// Endpoint untuk Categories
	categories := api.Group("/categories")
	categories.Get("/", categoryHandler.GetCategories)
	categories.Get("/:id", categoryHandler.GetCategory)
	categories.Post("/", categoryHandler.CreateCategory)
	categories.Put("/:id", categoryHandler.UpdateCategory)
	categories.Delete("/:id", categoryHandler.DeleteCategory)

	// Endpoint untuk Transactions (nested di dalam items)
	items.Get("/:id/transactions", transactionHandler.GetTransactions)
	items.Post("/:id/transactions", transactionHandler.CreateTransaction)

	log.Fatal(app.Listen(":8080"))
}
