package routes

import (
	"github.com/joaquinrs05/BuyRush/services/product-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/products", func(c *fiber.Ctx) error { return handlers.GetProducts(c, db) })
	app.Get("/products/:id", func(c *fiber.Ctx) error { return handlers.GetProductById(c, db) })
	app.Post("/products", func(c *fiber.Ctx) error { return handlers.CreateProduct(c, db) })
	app.Put("/products/:id", func(c *fiber.Ctx) error { return handlers.UpdateProduct(c, db) })
	app.Delete("/products/:id", func(c *fiber.Ctx) error { return handlers.DeleteProduct(c, db) })
}
