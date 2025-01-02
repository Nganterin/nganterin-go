package routers

import (
	"nganterin-go/middleware"
	"nganterin-go/reviews/controllers"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.RouterGroup, reviewControllers controllers.CompControllers) {
	reviewGroup := r.Group("/review")
	reviewGroup.Use(middleware.AuthMiddleware())
	{
		hotelGroup := reviewGroup.Group("/hotel")
		{
			hotelGroup.POST("/register", reviewControllers.Create)
		}
	}
}