package repositories

import (
	"nganterin-go/exceptions"
	"nganterin-go/models/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompRepositories interface {
	Create(ctx *gin.Context, tx *gorm.DB, data database.HotelReviews) *exceptions.Exception
	FindByHotelID(ctx *gin.Context, tx *gorm.DB, id string) ([]database.HotelReviews, *exceptions.Exception)
}