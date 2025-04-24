package db

import (
	"zadanie_go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	database.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Cart{},
		&models.User{},
		&models.Order{},
		&models.CartItem{},
	)

	DB = database
}
