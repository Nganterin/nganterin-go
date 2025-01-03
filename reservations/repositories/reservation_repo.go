package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data database.HotelReservations) *exceptions.Exception
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.HotelReservations, *exceptions.Exception) 
	FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelOrders, *exceptions.Exception)
	FindByReservationKey(ctx *gin.Context, tx *gorm.DB, reservationKey string) (*database.HotelOrders, *exceptions.Exception)
	CheckIn(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception
	CheckOut(ctx *gin.Context, tx *gorm.DB, reservationKey string) *exceptions.Exception
	Reviewed(ctx *gin.Context, tx *gorm.DB, orderID string) *exceptions.Exception
}
