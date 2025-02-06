package routers

import (
	"nganterin-go/api/reservations/controllers"

	"github.com/gin-gonic/gin"
)

func ReservationRoutes(r *gin.RouterGroup, reservationControllers controllers.CompControllers) {
	reservationGroup := r.Group("/reservation")
	{
		hotelGroup := reservationGroup.Group("/hotel")
		{
			hotelGroup.GET("/getall", reservationControllers.FindByHotelID)
			hotelGroup.GET("/details", reservationControllers.FindByReservationKey)
			hotelGroup.POST("/checkin", reservationControllers.CheckIn)
			hotelGroup.POST("/checkout", reservationControllers.CheckOut)
		}
	}
}
