package routers

import (
	"nganterin-go/api/hotels/controllers"

	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.RouterGroup, hotelControllers controllers.CompControllers) {
	hotelGroup := r.Group("/hotel")
	{
		hotelGroup.GET("/getall", hotelControllers.FindByPartnerID)
		hotelGroup.POST("/register", hotelControllers.Create)
	}
}
