package handlers

import (
	"github.com/joaquinrs05/BuyRush/services/order-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetOrders(c *fiber.Ctx, db *gorm.DB) error {
	var orders []models.Order
	db.Find(&orders)
	return c.JSON(orders)
}

func GetOrderById(c *fiber.Ctx, db *gorm.DB) error {
	var order models.Order
	id := c.Params("id")
	if err := db.First(&order, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}
	return c.JSON(order)
}

func CreateOrder(c *fiber.Ctx, db *gorm.DB) error {
	order := new(models.Order)
	if err := c.BodyParser(order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(order)
	return c.Status(fiber.StatusCreated).JSON(order)
}

func UpdateOrder(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var order models.Order
	if err := db.First(&order, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Order not found"})
	}
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Save(&order)
	return c.JSON(order)
}

func DeleteOrder(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	if err := db.Delete(&models.Order{}, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
