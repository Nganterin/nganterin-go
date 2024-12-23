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
	api.Use(middleware.GzipResponseMiddleware())

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

	hotelRoute := api.Group("/hotel")
	{
		hotelRoute.GET("/getall", compHandler.GetAllHotels)
		hotelRoute.GET("/details", compHandler.GetHotelByID)
		hotelRoute.GET("/search", compHandler.SearchHotels)
	}

	orderRoute := api.Group("/order")
	orderRoute.Use(middleware.AuthMiddleware())
	{
		hotelRoute := orderRoute.Group("/hotel")
		{
			hotelRoute.POST("/register", compHandler.RegisterHotelOrder)
			hotelRoute.GET("/get", compHandler.GetHotelOrderByID)
			hotelRoute.GET("/getall", compHandler.GetAllHotelOrderByUserID)
		}
	}

	filesRoute := api.Group("/files")
	{
		filesRoute.POST("/upload", compHandler.FileUpload)
	}

	partnerRoute := api.Group("/partner")
	{
		partnerAuthRoute := partnerRoute.Group("/auth")
		{
			partnerAuthRoute.POST("/register", compHandler.RegisterPartner)
			partnerAuthRoute.POST("/login", compHandler.LoginPartner)
			partnerAuthRoute.POST("/verify", compHandler.VerifyPartnerEmail)
		}

		partnerRoute.Use(middleware.PartnerAuthMiddleware())
		{
			hotelRoute := partnerRoute.Group("/hotel")
			{
				hotelRoute.POST("/register", compHandler.RegisterHotel)
			}
		}
	}

	midtransRoute := api.Group("/midtrans")
	{
		midtransRoute.POST("/notification", compHandler.MidtransNotification)
	}
}
