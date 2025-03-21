package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"orgmanager/models"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	database, err := gorm.Open(sqlite.Open("employees.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	database.AutoMigrate(&models.Employee{})
	DB = database
	return DB
}
