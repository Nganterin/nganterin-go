package routers

import (
	"nganterin-go/api/orders/controllers"
	"nganterin-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.RouterGroup, orderControllers controllers.CompControllers) {
	orderGroup := r.Group("/order")
	orderGroup.Use(middleware.AuthMiddleware())
	{
		hotelGroup := orderGroup.Group("/hotel")
		{
			hotelGroup.POST("/register", orderControllers.Create)
			hotelGroup.GET("/get", orderControllers.FindByID)
			hotelGroup.GET("/getall", orderControllers.FindByUserID)
		}
	}
}
