package handlers

import (
	"user-service/models"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_secret_key")

func GetUsers(c *fiber.Ctx, db *gorm.DB) error {
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

/*FIXME estas dos funciones probablemente sean iguales*/
func CreateUser(c *fiber.Ctx, db *gorm.DB) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func GetUserId(c *fiber.Ctx, db *gorm.DB) error {
	var user models.User
	id := c.Params("id")
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}
	return c.JSON(user)
}

func Register(c *fiber.Ctx, db *gorm.DB) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Aquí deberías hashear la contraseña antes de guardar
	db.Create(user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx, db *gorm.DB) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}
	// TODO: Compare hashed password here
	if user.Password != input.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Incorrect password"})
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // expires in 72 hours
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not generate token"})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
		"user":  user,
	})
}

func UpdateUser(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Usuario no encontrado"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	db.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx, db *gorm.DB) error {
	id := c.Params("id")
	if err := db.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
