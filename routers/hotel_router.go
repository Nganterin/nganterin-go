package routers

import (
	"nganterin-go/api/hotels/controllers"

	"github.com/gin-gonic/gin"
)

func HotelRoutes(r *gin.RouterGroup, hotelController controllers.CompControllers) {
	hotelGroup := r.Group("/hotel")
	{
		hotelGroup.GET("/getall", hotelController.FindAll)
		hotelGroup.GET("/details", hotelController.FindByID)
		hotelGroup.GET("/search", hotelController.SearchEngine)
	}
}
