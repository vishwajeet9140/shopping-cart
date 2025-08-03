package main

import (
	"github.com/gin-gonic/gin"
	"shopping-cart/database"
	"shopping-cart/routes"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	routes.SetupRoutes(r)
	r.Run(":8000")
}
