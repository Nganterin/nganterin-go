package routers

import (
	"nganterin-go/api/reservations/controllers"
	"nganterin-go/pkg/middleware"

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
