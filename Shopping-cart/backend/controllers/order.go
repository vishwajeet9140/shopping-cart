package controllers

import (
	"net/http"
	"shopping-cart/database"
	"shopping-cart/models"

	"github.com/gin-gonic/gin"
)

func PlaceOrder(c *gin.Context) {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var cart models.Cart
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	order := models.Order{
		CartID: cart.ID,
		UserID: userID,
	}

	database.DB.Create(&order)

	// Optional: clear cart items if needed
	database.DB.Where("cart_id = ?", cart.ID).Delete(&models.CartItem{})

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order_id": order.ID})
}

func GetOrders(c *gin.Context) {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var orders []models.Order
	database.DB.Where("user_id = ?", userID).Find(&orders)

	c.JSON(http.StatusOK, orders)
}
