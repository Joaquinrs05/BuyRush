package routes

import (
	"user-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/users", func(c *fiber.Ctx) error {
		return handlers.GetUsers(c, db)
	})

	app.Post("/users", func(c *fiber.Ctx) error {
		return handlers.CreateUser(c, db)
	})
}
