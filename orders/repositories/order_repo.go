package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data database.HotelOrders) *exceptions.Exception
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*database.HotelOrders, *exceptions.Exception)
	UpdatePaymentStatus(ctx *gin.Context, tx *gorm.DB, id string, status string) *exceptions.Exception
	FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelOrders, *exceptions.Exception)
}
