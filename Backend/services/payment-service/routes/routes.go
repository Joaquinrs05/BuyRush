package routes

import (
	"github.com/joaquinrs05/BuyRush/services/payment-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/payments", func(c *fiber.Ctx) error { return handlers.GetPayments(c, db) })
	app.Get("/payments/:id", func(c *fiber.Ctx) error { return handlers.GetPaymentById(c, db) })
	app.Post("/payments", func(c *fiber.Ctx) error { return handlers.CreatePayment(c, db) })
}
