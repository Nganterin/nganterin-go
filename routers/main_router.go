package routers

import (
	"github.com/gin-gonic/gin"
	"nganterin-go/config"
	"nganterin-go/handlers"
	"nganterin-go/middleware"
	"nganterin-go/repositories"
	"nganterin-go/services"
	hotelHandlers "nganterin-go/handlers/hotels"
	hotelRepos "nganterin-go/repositories/hotels"
	hotelServices "nganterin-go/services/hotels"
)

func CompRouter(api *gin.RouterGroup) {
	db := config.InitDB()

	api.Use(middleware.ClientTracker(db))

	compRepository := repositories.NewComponentRepository(db)
	hotelRepository := hotelRepos.NewHotelRepository(db)

	compService := services.NewService(compRepository)
	hotelService := hotelServices.NewHotelService(hotelRepository)

	compHandler := handlers.NewCompHandlers(compService)
	hotelHandler := hotelHandlers.NewHotelHandler(hotelService)

	api.GET("/ping", compHandler.Ping)

	authRoute := api.Group("/auth")
	{
		authRoute.POST("/register", compHandler.RegisterUserCredential)
		authRoute.POST("/login", compHandler.LoginUserCredentials)
		authRoute.POST("/verify", compHandler.VerifyUserEmail)
	}

	authRoute.Use(middleware.AuthMiddleware())
	{
		authRoute.GET("/auth-test", compHandler.AuthTest)
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
		hotelRoute.POST("/register", hotelHandler.CreateHotel)
	}
}
