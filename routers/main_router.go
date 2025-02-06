package routers

import (
	"net/http"
	"nganterin-go/injectors"
	"nganterin-go/models/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, dto.Response{
			Status: http.StatusOK,
			Data:   "pong",
		})
	})

	userController := injectors.InitializeUserController(db, validate)
	hotelController := injectors.InitializeHotelController(db, validate)
	orderController := injectors.InitializeOrderController(db, validate)
	storageController := injectors.InitializeStorageController(db, validate)
	partnerController := injectors.InitializePartnerController(db, validate)
	reservationController := injectors.InitializeReservationController(db, validate)
	reviewController := injectors.InitializeReviewController(db, validate)

	AuthRoutes(r, userController)
	HotelRoutes(r, hotelController)
	OrderRoutes(r, orderController)
	StorageRoutes(r, storageController)
	PartnerRoutes(r, partnerController, hotelController, reservationController, orderController)
	ReservationRoutes(r, reservationController)
	ReviewRoutes(r, reviewController)
}
