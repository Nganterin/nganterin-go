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
	db := config.InitDB()

	api.Use(middleware.ClientTracker(db))

	compRepository := repositories.NewComponentRepository(db)

	compService := services.NewService(compRepository)

	compHandler := handlers.NewCompHandlers(compService)

	api.GET("/ping", compHandler.Ping)

	authRoute := api.Group("/auth")
	{
		authRoute.POST("/register", compHandler.RegisterUserCredential)
		authRoute.POST("/login", compHandler.LoginUserCredentials)
		authRoute.POST("/verify", compHandler.VerifyUserEmail)

		authRoute.Use(middleware.AuthMiddleware())
		{
			authRoute.GET("/auth-test", compHandler.AuthTest)
		}
	}

	partnerRoute := api.Group("/partner")
	{
		partnerAuthRoute := partnerRoute.Group("/auth")
		{
			partnerAuthRoute.POST("/register", compHandler.RegisterPartner)
		}
	}

	hotelRoute := api.Group("/hotel")
	{
		hotelRoute.GET("/getall", compHandler.GetAllHotels)

		hotelRoute.Use(middleware.PartnerAuthMiddleware())
		{
			hotelRoute.POST("/register", compHandler.RegisterHotel)
		}
	}
}
