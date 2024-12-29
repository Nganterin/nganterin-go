package routers

import (
	"nganterin-go/middleware"
	"nganterin-go/users/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	authGroup := r.Group("/auth")
	{	
		authGroup.POST("/register", userController.CreateCredentials)
		authGroup.POST("/login", userController.LoginCredentials)
		authGroup.POST("/verify", userController.VerifyEmail)

		googleGroup := r.Group("/google")
		{
			googleGroup.POST("/login", userController.LoginGoogleOAuth)
		}

		authGroup.Use(middleware.AuthMiddleware())
		{
			authGroup.GET("/auth-test", userController.AuthTest)
		}
	}
}
