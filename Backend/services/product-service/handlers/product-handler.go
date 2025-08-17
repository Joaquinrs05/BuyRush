package handlers

import (
	"github.com/joaquinrs05/BuyRush/services/product-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetProducts(c *fiber.Ctx, db *gorm.DB) error {
	var products []models.Product
	db.Find(&products)
	return c.JSON(products)
}

func GetProductById(c *fiber.Ctx, db *gorm.DB) error {
	var product models.Product
	id := c.Params("id")
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func CreateProduct(c *fiber.Ctx, db *gorm.DB) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(product)
	return c.Status(fiber.StatusCreated).JSON(product)
}

func UpdateProduct(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var product models.Product
	if err := db.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Save(&product)
	return c.JSON(product)
}

func DeleteProduct(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	if err := db.Delete(&models.Product{}, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
