package routers

import (
	publicInjector "nganterin-go/injectors"
	"nganterin-go/partners/injectors"
	"nganterin-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PartnerRoutes(r *gin.RouterGroup, db *gorm.DB, validate *validator.Validate) {
	partnerController := injectors.InitializePartnerController(db, validate)

	hotelController := publicInjector.InitializeHotelController(db, validate)
	orderController := publicInjector.InitializeOrderController(db, validate)
	reservationController := publicInjector.InitializeReservationController(db, validate)

	AuthRoutes(r, partnerController)
	r.Use(middleware.PartnerAuthMiddleware())

	AnalyticRoutes(r, reservationController, orderController)
	ApprovalRoutes(r, partnerController)
	HotelRoutes(r, hotelController)
	ReservationRoutes(r, reservationController)
}
