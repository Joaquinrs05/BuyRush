package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

func jwtMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or invalid token"})
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
	}
	return c.Next()
}

func main() {
	app := fiber.New()

	// Public routes (no JWT required)
	app.Post("/register", proxy.Forward("http://localhost:8081/register"))
	app.Post("/login", proxy.Forward("http://localhost:8081/login"))

	// JWT Middleware for protected routes
	app.Use(jwtMiddleware)

	// Protected routes (require JWT)
	app.All("/users*", proxy.Forward("http://localhost:8081"))
	app.All("/products*", proxy.Forward("http://localhost:8082"))
	app.All("/orders*", proxy.Forward("http://localhost:8083"))
	app.All("/payments*", proxy.Forward("http://localhost:8084"))
	app.All("/notifications*", proxy.Forward("http://localhost:8085"))

	app.Listen(":8080")
}
