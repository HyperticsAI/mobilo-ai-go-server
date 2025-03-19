package main

import (
	"fmt"
	"go-server/config"
	"go-server/models"
	"go-server/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbConfig := config.LoadDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.DBHost,
		dbConfig.DBUser,
		dbConfig.DBPassword,
		dbConfig.DBName,
		dbConfig.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	models.DB = db

	// AutoMigrate all models
	db.AutoMigrate(&models.OrganizationSetting{}, &models.AIResponse{})

	r := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	r.Run(":" + port)
}
