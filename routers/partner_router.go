package routers

import (
	"nganterin-go/middleware"
	"nganterin-go/partners/controllers"

	hotelControllers "nganterin-go/hotels/controllers"
	orderControllers "nganterin-go/orders/controllers"
	reservationControllers "nganterin-go/reservations/controllers"

	"github.com/gin-gonic/gin"
)

func PartnerRoutes(
	r *gin.RouterGroup,
	partnerControllers controllers.CompControllers,
	hotelControllers hotelControllers.CompControllers,
	reservationControllers reservationControllers.CompControllers,
	orderControllers orderControllers.CompControllers,
) {
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
			hotelGroup := partnerGroup.Group("/hotel")
			{
				hotelGroup.GET("/getall", hotelControllers.FindByPartnerID)
				hotelGroup.POST("/register", hotelControllers.Create)
			}

			reservationGroup := partnerGroup.Group("/reservation")
			{
				hotelGroup := reservationGroup.Group("/hotel")
				{
					hotelGroup.GET("/details", reservationControllers.FindByReservationKey)
					hotelGroup.POST("/checkin", reservationControllers.CheckIn)
					hotelGroup.POST("/checkout", reservationControllers.CheckOut)
				}
			}

			analyticGroup := partnerGroup.Group("/analytic")
			{
				reservationGroup := analyticGroup.Group("/reservation")
				{
					reservationGroup.GET("/yearly", reservationControllers.YearlyReservationAnalytic)
				}

				orderGroup := analyticGroup.Group("/order")
				{
					orderGroup.GET("/yearly", orderControllers.YearlyOrderAnalytic)
				}
			}

			approvalGroup := partnerGroup.Group("/approval")
			{
				approvalGroup.GET("/status", partnerControllers.ApprovalCheck)
			}
		}
	}
}
