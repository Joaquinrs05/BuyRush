package routes

import (
	"github.com/joaquinrs05/BuyRush/services/notification-service/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/notifications", func(c *fiber.Ctx) error { return handlers.GetNotifications(c, db) })
	app.Get("/notifications/:id", func(c *fiber.Ctx) error { return handlers.GetNotificationById(c, db) })
	app.Post("/notifications", func(c *fiber.Ctx) error { return handlers.CreateNotification(c, db) })
	app.Put("/notifications/:id/read", func(c *fiber.Ctx) error { return handlers.MarkAsRead(c, db) })
}
