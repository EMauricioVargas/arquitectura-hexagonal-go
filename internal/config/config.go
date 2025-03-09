package config

import (
	"fmt"
	"github.com/EMauricioVargas/arquitectura-hexagonal-go/internal/domain/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func LoadConfig() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Auto migración de modelos
	err = DB.AutoMigrate(&entities.Product{}, &entities.User{}, &entities.Bill{}, &entities.BillDetail{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
