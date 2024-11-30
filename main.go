package main

import (
	"order-management/database"
	"order-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()

	r := gin.Default()

	routes.AuthRoutes(r)
	routes.OrderRoutes(r)

	r.Run(":8080")
}
