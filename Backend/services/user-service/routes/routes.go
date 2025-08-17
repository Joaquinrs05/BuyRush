package routes

import (
	"user-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
    app.Post("/register", func(c *fiber.Ctx) error {
        return handlers.Register(c, db)
    })

    app.Post("/login", func(c *fiber.Ctx) error {
        return handlers.Login(c, db)
    })

    app.Get("/users", func(c *fiber.Ctx) error {
        return handlers.GetUsers(c, db)
    })

    app.Get("/users/:id", func(c *fiber.Ctx) error {
        return handlers.GetUserId(c, db)
    })

    app.Put("/users/:id", func(c *fiber.Ctx) error {
        return handlers.UpdateUser(c, db)
    })

    app.Delete("/users/:id", func(c *fiber.Ctx) error {
        return handlers.DeleteUser(c, db)
    })
}