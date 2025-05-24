package db

import (
	"backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = database.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.Cart{},
		&models.User{},
		&models.Order{},
		&models.CartItem{},
	)
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	DB = database

	var count int64
	DB.Model(&models.Category{}).Count(&count)
	if count == 0 {
		DB.Create(&models.Category{Name: "Test category"})
	}
	DB.Model(&models.Product{}).Count(&count)
	if count == 0 {
		DB.Create(&models.Product{Name: "Test Product", Price: 9.99, CategoryID: 1})
	}
}
