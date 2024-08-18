package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gofiber-templ-bench/pkg/models"
	"gofiber-templ-bench/pkg/users"
	"gofiber-templ-bench/utils/i18n"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

func main() {
	app := fiber.New()

	slog.Info("Using local db ./test.db")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	slog.Info("Migrating Database")
	err = db.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		panic("migration failed: " + err.Error())
	}

	// Add logging & i18n to all requests
	app.Use(logger.New())
	app.Use(i18n.NewMiddleware())

	// Setup user routes
	usersHandler := users.NewHandler(db)
	app.Get("/", usersHandler.Index)
	app.Get("/user/:id<int>/edit", usersHandler.Edit)
	app.Post("/user/:id<int>", usersHandler.Update)

	// seed
	slog.Info("Seeding Database")
	address1 := models.Address{Name: "Street 1, 1234 City"}
	if err := db.Create(&address1).Error; err != nil {
		panic("failed to seed address: " + err.Error())
	}
	address2 := models.Address{Name: "Street 2, 4321 Ytic"}
	if err := db.Create(&address2).Error; err != nil {
		panic("failed to seed address: " + err.Error())
	}

	user := models.User{Name: "Ellen Doe", Email: "ellen@example.com", AddressID: address1.ID}
	if err := db.Create(&user).Error; err != nil {
		panic("failed to seed user: " + err.Error())
	}

	slog.Info("Starting server on port 3000")
	log.Fatal(app.Listen(":3000"))
}
