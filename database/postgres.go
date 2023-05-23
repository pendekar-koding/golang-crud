package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rest-api-crud/models"
)

var DB *gorm.DB
var err error

func DatabaseConnection() {
	host := "localhost"
	port := "5432"
	dbName := "golang-crud"
	dbUser := "postgres"
	password := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		dbUser,
		password,
		dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err := DB.AutoMigrate(models.Book{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")
}
