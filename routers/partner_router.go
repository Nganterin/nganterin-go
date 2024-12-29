package routers

import (
	hotelControllers "nganterin-go/hotels/controllers"
	"nganterin-go/middleware"
	"nganterin-go/partners/controllers"

	"github.com/gin-gonic/gin"
)

func PartnerRoutes(r *gin.RouterGroup, partnerControllers controllers.CompControllers, hotelControllers hotelControllers.CompControllers) {
	partnerGroup := r.Group("/partner")
	{
		partnerAuthGroup := partnerGroup.Group("/auth")
		{
			partnerAuthGroup.POST("/register", partnerControllers.Create)
			partnerAuthGroup.POST("/login", partnerControllers.Login)
			partnerAuthGroup.POST("/verify", partnerControllers.VerifyEmail)
		}

		partnerGroup.Use(middleware.PartnerAuthMiddleware())
		{
			hotelRoute := partnerGroup.Group("/hotel")
			{
				hotelRoute.POST("/register", hotelControllers.Create)
			}
		}
	}
}
