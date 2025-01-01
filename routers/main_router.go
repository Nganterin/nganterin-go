package routers

import (
	"net/http"
	"nganterin-go/config"
	"nganterin-go/injectors"
	"nganterin-go/middleware"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
)

func CompRouters(api *gin.RouterGroup) {
	db := config.InitDB()

	api.Use(middleware.ClientTracker(db))
	api.Use(middleware.GzipResponseMiddleware())

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.Response{
			Status: http.StatusOK,
			Data:   "pong",
		})
	})

	userController := injectors.InitializeUserController(db)
	hotelController := injectors.InitializeHotelController(db)
	orderController := injectors.InitializeOrderController(db)
	midtransController := injectors.InitializeMidtransController(db)
	storageController := injectors.InitializeStorageController(db)
	partnerController := injectors.InitializePartnerController(db)
	reservationController := injectors.InitializeReservationController(db)

	AuthRoutes(api, userController)
	HotelRoutes(api, hotelController)
	OrderRoutes(api, orderController)
	MidtransRoutes(api, midtransController)
	StorageRoutes(api, storageController)
	PartnerRoutes(api, partnerController, hotelController, reservationController)
	ReservationRoutes(api, reservationController)
}
