package routers

import (
	"nganterin-go/midtrans/notifications/controllers"

	"github.com/gin-gonic/gin"
)

func MidtransRoutes(r *gin.RouterGroup, midtransController controllers.CompControllers) {
	notificationGroup := r.Group("/notification")
	{
		notificationGroup.POST("/payment", midtransController.Notification)
	}
}
