package repositories

import (
	"nganterin-go/models"
	"nganterin-go/pkg/exceptions"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data models.HotelReviews) *exceptions.Exception
	FindByHotelID(ctx *gin.Context, tx *gorm.DB, id string) ([]models.HotelReviews, *exceptions.Exception)
}
