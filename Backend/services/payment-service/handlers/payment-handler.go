package handlers

import (
	"github.com/joaquinrs05/BuyRush/services/payment-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetPayments(c *fiber.Ctx, db *gorm.DB) error {
	var payments []models.Payment
	db.Find(&payments)
	return c.JSON(payments)
}

func GetPaymentById(c *fiber.Ctx, db *gorm.DB) error {
	var payment models.Payment
	id := c.Params("id")
	if err := db.First(&payment, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Payment not found"})
	}
	return c.JSON(payment)
}

func CreatePayment(c *fiber.Ctx, db *gorm.DB) error {
	payment := new(models.Payment)
	if err := c.BodyParser(payment); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(payment)
	return c.Status(fiber.StatusCreated).JSON(payment)
}
