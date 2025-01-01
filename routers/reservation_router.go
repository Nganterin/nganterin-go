package routers

import (
	"nganterin-go/middleware"
	"nganterin-go/reservations/controllers"

	"github.com/gin-gonic/gin"
)

func ReservationRoutes(r *gin.RouterGroup, reservationControllers controllers.CompControllers) {
	reservationGroup := r.Group("/reservation")
	reservationGroup.Use(middleware.AuthMiddleware())
	{
		hotelGroup := reservationGroup.Group("/hotel")
		{
			hotelGroup.GET("/getall", reservationControllers.FindByUserID)
		}
	}
}