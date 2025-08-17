package routes

import (
	"github.com/joaquinrs05/BuyRush/services/order-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/orders", func(c *fiber.Ctx) error { return handlers.GetOrders(c, db) })
	app.Get("/orders/:id", func(c *fiber.Ctx) error { return handlers.GetOrderById(c, db) })
	app.Post("/orders", func(c *fiber.Ctx) error { return handlers.CreateOrder(c, db) })
	app.Put("/orders/:id", func(c *fiber.Ctx) error { return handlers.UpdateOrder(c, db) })
	app.Delete("/orders/:id", func(c *fiber.Ctx) error { return handlers.DeleteOrder(c, db) })
}
