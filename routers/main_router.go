package routers

import (
	"nganterin-go/config"
	"nganterin-go/handlers"
	"nganterin-go/middleware"
	"nganterin-go/repositories"
	"nganterin-go/services"

	"github.com/gin-gonic/gin"
)

func CompRouter(api *gin.RouterGroup) {
	api.Use(middleware.ClientTracker(config.InitDB()))

	compRepository := repositories.NewComponentRepository(config.InitDB())
	compService := services.NewService(compRepository)
	compHandler := handlers.NewCompHandlers(compService)

	api.GET("/ping", compHandler.Ping)

	authRoute := api.Group("/auth")
	{
		authRoute.POST("/register", compHandler.RegisterUserCredential)
		authRoute.POST("/login", compHandler.LoginUserCredentials)
	}

	authRoute.Use(middleware.AuthMiddleware())
	{
		authRoute.GET("/auth-test", compHandler.AuthTest)
	}

}
