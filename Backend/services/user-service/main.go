package main

import (
	"user-service/config"
	"user-service/database"
	"user-service/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	db := database.Connect()
	database.Migrate(db)

	app := fiber.New()
	routes.RegisterRoutes(app, db)

	app.Listen(":8080")
}
