package repositories

import (
	"nganterin-go/models"
	"nganterin-go/models/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.HotelReservations) *exceptions.Exception
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*models.HotelReservations, *exceptions.Exception)
	FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]models.HotelOrders, *exceptions.Exception)
	FindByHotelID(ctx *gin.Context, tx *gorm.DB, hotelID string) ([]models.HotelOrders, *exceptions.Exception)
	FindByReservationKey(ctx *gin.Context, tx *gorm.DB, reservationKey string) (*models.HotelOrders, *exceptions.Exception)
	CheckIn(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception
	CheckOut(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception
	Reviewed(ctx *gin.Context, tx *gorm.DB, orderID string) *exceptions.Exception
	FindLast12MonthReservationCount(ctx *gin.Context, tx *gorm.DB, partnerID string) ([]dto.HotelMonthlyReservation, *exceptions.Exception)
}
