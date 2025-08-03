package routes

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/controllers"
	"shopping-cart/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.GetUsers)

	r.POST("/items", controllers.CreateItem)
	r.GET("/items", controllers.GetItems)

	protected := r.Group("/", middleware.AuthMiddleware())
	{
		protected.POST("/carts", controllers.AddToCart)
		protected.GET("/carts", controllers.GetCarts)

		protected.POST("/orders", controllers.PlaceOrder)
		protected.GET("/orders", controllers.GetOrders)
	}
}
