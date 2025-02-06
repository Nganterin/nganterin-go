package routers

import (
	orderControllers "nganterin-go/api/orders/controllers"
	reservationController "nganterin-go/api/reservations/controllers"

	"github.com/gin-gonic/gin"
)

func AnalyticRoutes(r *gin.RouterGroup, reservationControllers reservationController.CompControllers, orderControllers orderControllers.CompControllers) {
	analyticGroup := r.Group("/analytic")
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
}
