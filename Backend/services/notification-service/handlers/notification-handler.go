package handlers

import (
	"github.com/joaquinrs05/BuyRush/services/notification-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetNotifications(c *fiber.Ctx, db *gorm.DB) error {
	var notifications []models.Notification
	db.Find(&notifications)
	return c.JSON(notifications)
}

func GetNotificationById(c *fiber.Ctx, db *gorm.DB) error {
	var notification models.Notification
	id := c.Params("id")
	if err := db.First(&notification, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Notification not found"})
	}
	return c.JSON(notification)
}

func CreateNotification(c *fiber.Ctx, db *gorm.DB) error {
	notification := new(models.Notification)
	if err := c.BodyParser(notification); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(notification)
	return c.Status(fiber.StatusCreated).JSON(notification)
}

func MarkAsRead(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var notification models.Notification
	if err := db.First(&notification, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Notification not found"})
	}
	notification.Read = true
	db.Save(&notification)
	return c.JSON(notification)
}
