package routers

import (
	"net/http"
	"nganterin-go/config"
	"nganterin-go/injectors"
	"nganterin-go/middleware"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CompRouters(api *gin.RouterGroup) {
	db := config.InitDB()
	validate := validator.New(validator.WithRequiredStructEnabled())

	api.Use(middleware.ClientTracker(db))
	api.Use(middleware.GzipResponseMiddleware())

	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.Response{
			Status: http.StatusOK,
			Data:   "pong",
		})
	})

	userController := injectors.InitializeUserController(db, validate)
	hotelController := injectors.InitializeHotelController(db, validate)
	orderController := injectors.InitializeOrderController(db, validate)
	midtransController := injectors.InitializeMidtransController(db, validate)
	storageController := injectors.InitializeStorageController(db, validate)
	partnerController := injectors.InitializePartnerController(db, validate)
	reservationController := injectors.InitializeReservationController(db, validate)
	reviewController := injectors.InitializeReviewController(db, validate)

	AuthRoutes(api, userController)
	HotelRoutes(api, hotelController)
	OrderRoutes(api, orderController)
	MidtransRoutes(api, midtransController)
	StorageRoutes(api, storageController)
	PartnerRoutes(api, partnerController, hotelController, reservationController)
	ReservationRoutes(api, reservationController)
	ReviewRoutes(api, reviewController)
}