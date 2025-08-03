package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"shopping-cart/models"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB!")
	}

	database.AutoMigrate(&models.User{}, &models.Item{}, &models.Cart{}, &models.CartItem{}, &models.Order{})
	DB = database
}
