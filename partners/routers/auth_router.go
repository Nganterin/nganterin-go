package routers

import (
	"nganterin-go/partners/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, partnerControllers controllers.CompControllers ) {
	partnerAuthGroup := r.Group("/auth")
	{
		partnerAuthGroup.POST("/register", partnerControllers.Create)
		partnerAuthGroup.POST("/login", partnerControllers.Login)
		partnerAuthGroup.POST("/verify", partnerControllers.VerifyEmail)
	}
}
