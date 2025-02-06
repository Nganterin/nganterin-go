package routers

import (
	"nganterin-go/midtrans/notifications/controllers"

	"github.com/gin-gonic/gin"
)

func MidtransRoutes(r *gin.RouterGroup, midtransController controllers.CompControllers) {
	midtransGroup := r.Group("/midtrans")
	{
		midtransGroup.POST("/notification", midtransController.Notification)
	}
}
