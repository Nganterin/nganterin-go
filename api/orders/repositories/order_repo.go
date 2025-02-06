package repositories

import (
	"nganterin-go/models"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.HotelOrders) *exceptions.Exception
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*models.HotelOrders, *exceptions.Exception)
	UpdatePaymentStatus(ctx *gin.Context, tx *gorm.DB, id string, status string) *exceptions.Exception
	FindByUserID(ctx *gin.Context, tx *gorm.DB, id string) ([]models.HotelOrders, *exceptions.Exception)
	FindByPartnerID(ctx *gin.Context, tx *gorm.DB, partnerID string) ([]models.HotelOrders, *exceptions.Exception)
}
