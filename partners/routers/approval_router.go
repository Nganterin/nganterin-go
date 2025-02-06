package routers

import (
	"nganterin-go/partners/controllers"

	"github.com/gin-gonic/gin"
)

func ApprovalRoutes(r *gin.RouterGroup, partnerControllers controllers.CompControllers) {
	approvalGroup := r.Group("/approval")
	{
		approvalGroup.GET("/status", partnerControllers.ApprovalCheck)
	}
}
