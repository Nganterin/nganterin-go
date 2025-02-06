package routers

import (
	"net/http"
	"nganterin-go/injectors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func CompRouters(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "pong",
		})
	})

	userController := injectors.InitializeUserController(db, validate)
	hotelController := injectors.InitializeHotelController(db, validate)
	orderController := injectors.InitializeOrderController(db, validate)
	storageController := injectors.InitializeStorageController(db, validate)
	reservationController := injectors.InitializeReservationController(db, validate)
	reviewController := injectors.InitializeReviewController(db, validate)

	AuthRoutes(r, userController)
	HotelRoutes(r, hotelController)
	OrderRoutes(r, orderController)
	StorageRoutes(r, storageController)
	ReservationRoutes(r, reservationController)
	ReviewRoutes(r, reviewController)
}
