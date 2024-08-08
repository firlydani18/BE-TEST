package main

import (
	"account-manager/models"
	"account-manager/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Connect to the database
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Run migrations
	err = db.AutoMigrate(&models.User{}, &models.Account{}, &models.PaymentHistory{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	routes.SetupRoutes(r, db)

	log.Fatal(r.Run(":3001"))
}
