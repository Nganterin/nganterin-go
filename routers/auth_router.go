package routers

import (
	"nganterin-go/api/users/controllers"
	"nganterin-go/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup, userController controllers.CompControllers) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", userController.CreateCredentials)
		authGroup.POST("/login", userController.LoginCredentials)
		authGroup.POST("/verify", userController.VerifyEmail)

		googleGroup := authGroup.Group("/google")
		{
			googleGroup.POST("/login", userController.LoginGoogleOAuth)
			googleGroup.POST("/register", userController.CreateGoogleOAuth)
		}

		authGroup.Use(middleware.AuthMiddleware())
		{
			authGroup.GET("/auth-test", userController.AuthTest)
		}
	}
}
