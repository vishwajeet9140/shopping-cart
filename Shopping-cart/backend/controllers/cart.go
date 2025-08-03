package controllers

import (
	"net/http"
	"shopping-cart/database"
	"shopping-cart/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret")

func getUserIDFromToken(c *gin.Context) (uint, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := uint(claims["user_id"].(float64))
		return userID, nil
	} else {
		return 0, err
	}
}

func AddToCart(c *gin.Context) {
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	var request struct {
		ItemID uint `json:"item_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	database.DB.Where("user_id = ?", userID).FirstOrCreate(&cart, models.Cart{UserID: userID})

	cartItem := models.CartItem{CartID: cart.ID, ItemID: request.ItemID}
	database.DB.Create(&cartItem)

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart"})
}

func GetCarts(c *gin.Context) {
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

	var cartItems []models.CartItem
	database.DB.Where("cart_id = ?", cart.ID).Find(&cartItems)

	c.JSON(http.StatusOK, cartItems)
}
