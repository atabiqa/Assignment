package config

import (
	"assignment2/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbPort   = "5432"
	dbName   = "assignment-2"
	db       *gorm.DB
	err      error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("Failed Connect to Database")
	}
	defer fmt.Println("Successfully Connected to Database")

	db.AutoMigrate(models.Order{}, models.Item{})
	return db
}

func GetDB() *gorm.DB {
	return db
}
