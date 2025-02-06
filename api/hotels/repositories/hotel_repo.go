package repositories

import (
	"nganterin-go/models"
	"nganterin-go/api/hotels/dto"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.Hotels) (*string, *exceptions.Exception)
	FindAll(ctx *gin.Context, tx *gorm.DB) ([]models.Hotels, *exceptions.Exception)
	SearchEngine(ctx *gin.Context, tx *gorm.DB, searchInput dto.HotelSearch) ([]models.Hotels, *exceptions.Exception)
	FindByID(ctx *gin.Context, tx *gorm.DB, id string) (*models.Hotels, *exceptions.Exception)
	FindRoomByID(ctx *gin.Context, tx *gorm.DB, id uint) (*models.HotelRooms, *exceptions.Exception)
	FindByPartnerID(ctx *gin.Context, tx *gorm.DB, partnerID string) ([]models.Hotels, *exceptions.Exception)
}
