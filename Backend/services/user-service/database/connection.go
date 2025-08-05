package database

import (
	"fmt"
	"os"

	/*"gorm.io/driver/postgres"*/
	"gorm.io/gorm"
	"user-service/models"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
