package routes

import (
	"order-management/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/api/v1")
	{
		auth.POST("/login", controllers.Login)
	}
}
