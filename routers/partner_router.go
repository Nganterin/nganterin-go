package routers

import (
	"nganterin-go/middleware"
	"nganterin-go/partners/controllers"

	hotelControllers "nganterin-go/hotels/controllers"
	reservationControllers "nganterin-go/reservations/controllers"

	"github.com/gin-gonic/gin"
)

func PartnerRoutes(r *gin.RouterGroup, partnerControllers controllers.CompControllers, hotelControllers hotelControllers.CompControllers, reservationControllers reservationControllers.CompControllers) {
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

			reservationGroup := partnerGroup.Group("/reservation")
			{
				hotelRoute := reservationGroup.Group("/hotel")
				{
					hotelRoute.GET("/details", reservationControllers.FindByReservationKey)
					hotelRoute.POST("/checkin", reservationControllers.CheckIn)
					hotelRoute.POST("/checkout", reservationControllers.CheckOut)
				}
			}
		}
	}
}
