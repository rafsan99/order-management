package routes

import (
	"order-management/controllers"
	"order-management/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	orders := r.Group("/api/v1/orders")
	orders.Use(middlewares.AuthMiddleware)
	{
		orders.POST("", controllers.CreateOrder)
		orders.GET("/all", controllers.OrdersList)
	}
}
