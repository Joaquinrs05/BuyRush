package handlers

import (
	"user-service/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUsers(c *fiber.Ctx, db *gorm.DB) error {
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx, db *gorm.DB) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(user)
	return c.Status(fiber.StatusCreated).JSON(user)
}
